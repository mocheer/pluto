package axios

import (
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Cache 定义缓存实现的接口
// 用户可以实现此接口以支持自定义存储后端（Redis、文件等）
type Cache interface {
	// Get 根据键获取缓存的响应，如果未找到或已过期则返回 nil
	Get(key string) *CacheEntry

	// Set 将响应以指定的键和 TTL 存储到缓存中
	Set(key string, entry *CacheEntry, ttl time.Duration)

	// Delete 从缓存中删除一个条目
	Delete(key string)

	// Clear 清除缓存中的所有条目
	Clear()

	// Stats 返回缓存统计信息
	Stats() CacheStats
}

// CacheEntry 表示一个缓存的 HTTP 响应
type CacheEntry struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
	CreatedAt  time.Time
	ExpiresAt  time.Time
}

// IsExpired 检查缓存条目是否已过期
func (e *CacheEntry) IsExpired() bool {
	return time.Now().After(e.ExpiresAt)
}

// CacheStats 提供缓存统计信息
type CacheStats struct {
	Hits   int64
	Misses int64
	Size   int64
}

// CacheKeyFunc 是用于生成缓存键的函数类型
type CacheKeyFunc func(method MethodType, fullURL string, headers map[string]string) string

// CacheConfig 保存客户端的全局缓存配置
type CacheConfig struct {
	// Cache 是使用的缓存实现
	Cache Cache

	// DefaultTTL 是未指定每次请求 TTL 时的默认缓存 TTL
	// 值为 0 表示没有默认 TTL（必须在每次请求中指定）
	DefaultTTL time.Duration

	// KeyFunc 是生成缓存键的自定义函数
	// 如果为 nil，则使用默认键函数（方法 + URL）
	KeyFunc CacheKeyFunc

	// CacheableMethods 定义哪些 HTTP 方法可以被缓存
	// 如果为 nil，默认为 []string{"GET"}
	CacheableMethods []MethodType
}

// RequestCacheOptions 包含每次请求的缓存配置
type RequestCacheOptions struct {
	// Enabled 显式启用/禁用此请求的缓存
	// 如果为 nil，则禁用缓存（选择加入模型）
	Enabled *bool

	// TTL 设置此特定请求的 TTL
	// 如果设置，将覆盖全局 DefaultTTL
	TTL time.Duration

	// ForceRefresh 绕过缓存并获取最新数据
	// 新的响应仍会被缓存
	ForceRefresh bool

	// CustomKey 允许覆盖此请求的缓存键
	CustomKey string
}

// Bool 是一个辅助函数，用于创建 bool 值的指针
func Bool(v bool) *bool {
	return &v
}

// CacheEnabled 返回启用缓存的 RequestCacheOptions
func CacheEnabled(ttl time.Duration) *RequestCacheOptions {
	enabled := true
	return &RequestCacheOptions{
		Enabled: &enabled,
		TTL:     ttl,
	}
}

// CacheDisabled 返回禁用缓存的 RequestCacheOptions
func CacheDisabled() *RequestCacheOptions {
	enabled := false
	return &RequestCacheOptions{
		Enabled: &enabled,
	}
}

// DefaultCacheKeyFunc 根据方法和完整 URL 生成缓存键
func DefaultCacheKeyFunc(method MethodType, fullURL string, _ map[string]string) string {
	return string(method) + ":" + fullURL
}

// MemoryCache 是线程安全的内存缓存实现
type MemoryCache struct {
	entries         map[string]*memoryCacheEntry
	mu              sync.RWMutex
	hits            int64
	misses          int64
	maxSize         int
	cleanupInterval time.Duration
	stopChan        chan struct{}
	stopped         bool
}

type memoryCacheEntry struct {
	entry *CacheEntry
}

// MemoryCacheOptions 配置 MemoryCache
type MemoryCacheOptions struct {
	// MaxSize 是最大条目数（0 表示无限制）
	MaxSize int

	// CleanupInterval 是清理过期条目的间隔时间（默认：5 分钟）
	CleanupInterval time.Duration
}

// NewMemoryCache 创建一个新的内存缓存
func NewMemoryCache(opts *MemoryCacheOptions) *MemoryCache {
	if opts == nil {
		opts = &MemoryCacheOptions{}
	}

	cleanupInterval := opts.CleanupInterval
	if cleanupInterval == 0 {
		cleanupInterval = 5 * time.Minute
	}

	mc := &MemoryCache{
		entries:         make(map[string]*memoryCacheEntry),
		maxSize:         opts.MaxSize,
		cleanupInterval: cleanupInterval,
		stopChan:        make(chan struct{}),
	}

	// 启动后台清理协程
	go mc.cleanupLoop()

	return mc
}

// Get 根据键获取缓存的响应
func (c *MemoryCache) Get(key string) *CacheEntry {
	c.mu.RLock()
	entry, exists := c.entries[key]
	if !exists {
		c.mu.RUnlock()
		atomic.AddInt64(&c.misses, 1)
		return nil
	}

	// 在持有锁时复制条目以避免竞态条件
	entryCopy := &CacheEntry{
		Body:       entry.entry.Body,
		StatusCode: entry.entry.StatusCode,
		Headers:    entry.entry.Headers,
		CreatedAt:  entry.entry.CreatedAt,
		ExpiresAt:  entry.entry.ExpiresAt,
	}
	c.mu.RUnlock()

	if entryCopy.IsExpired() {
		// 删除过期条目
		c.Delete(key)
		atomic.AddInt64(&c.misses, 1)
		return nil
	}

	atomic.AddInt64(&c.hits, 1)
	return entryCopy
}

// Set 将响应存储到缓存中
func (c *MemoryCache) Set(key string, entry *CacheEntry, ttl time.Duration) {
	if ttl <= 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// 检查最大大小，必要时淘汰最旧的条目
	if c.maxSize > 0 && len(c.entries) >= c.maxSize {
		// 简单淘汰：删除第一个过期或最旧的条目
		c.evictOne()
	}

	entry.ExpiresAt = time.Now().Add(ttl)
	c.entries[key] = &memoryCacheEntry{
		entry: entry,
	}
}

// Delete 从缓存中删除一个条目
func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, key)
}

// Clear 清除缓存中的所有条目
func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]*memoryCacheEntry)
}

// Stats 返回缓存统计信息
func (c *MemoryCache) Stats() CacheStats {
	c.mu.RLock()
	size := int64(len(c.entries))
	c.mu.RUnlock()

	return CacheStats{
		Hits:   atomic.LoadInt64(&c.hits),
		Misses: atomic.LoadInt64(&c.misses),
		Size:   size,
	}
}

// Close 停止后台清理协程
func (c *MemoryCache) Close() {
	c.mu.Lock()
	if !c.stopped {
		c.stopped = true
		close(c.stopChan)
	}
	c.mu.Unlock()
}

// evictOne 移除一个条目以腾出空间（调用时已持有锁）
func (c *MemoryCache) evictOne() {
	var oldestKey string
	var oldestTime time.Time
	first := true

	for key, entry := range c.entries {
		// 首先尝试淘汰过期条目
		if entry.entry.IsExpired() {
			delete(c.entries, key)
			return
		}

		// 跟踪最旧的条目
		if first || entry.entry.CreatedAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.entry.CreatedAt
			first = false
		}
	}

	// 如果没有过期条目，则淘汰最旧的
	if oldestKey != "" {
		delete(c.entries, oldestKey)
	}
}

// cleanupLoop 定期运行以移除过期条目
func (c *MemoryCache) cleanupLoop() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanupExpired()
		case <-c.stopChan:
			return
		}
	}
}

// cleanupExpired 移除所有过期条目
func (c *MemoryCache) cleanupExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if entry.entry.IsExpired() {
			delete(c.entries, key)
		}
	}
}

// shouldCacheRequest 判断请求是否应该被缓存
func shouldCacheRequest(cacheOptions *CacheConfig, options *AxiosOptions) bool {
	// 检查客户端是否配置了缓存
	if cacheOptions == nil || cacheOptions.Cache == nil {
		return false
	}

	// 检查每次请求是否显式禁用了缓存
	if options.Cache != nil && options.Cache.Enabled != nil && !*options.Cache.Enabled {
		return false
	}

	// 检查每次请求是否显式启用了缓存
	if options.Cache != nil && options.Cache.Enabled != nil && *options.Cache.Enabled {
		return isMethodCacheable(cacheOptions, options.Method)
	}

	// 默认：除非每次请求显式启用，否则禁用缓存
	return false
}

// isMethodCacheable 检查 HTTP 方法是否可以被缓存
func isMethodCacheable(cacheConfig *CacheConfig, method MethodType) bool {
	methods := cacheConfig.CacheableMethods
	if methods == nil {
		methods = []MethodType{MethodGet}
	}
	for _, m := range methods {
		if strings.EqualFold(string(m), string(method)) {
			return true
		}
	}
	return false
}

// shouldForceRefresh 检查请求是否应绕过缓存
func shouldForceRefresh(options *AxiosOptions) bool {
	return options.Cache != nil && options.Cache.ForceRefresh
}

// generateCacheKey 为请求生成缓存键
func generateCacheKey(cacheConfig *CacheConfig, options *AxiosOptions, fullURL string) string {
	// 检查请求选项中是否有自定义键
	if options.Cache != nil && options.Cache.CustomKey != "" {
		return options.Cache.CustomKey
	}

	// 如果配置了自定义键函数则使用
	if cacheConfig.KeyFunc != nil {
		return cacheConfig.KeyFunc(options.Method, fullURL, options.Headers)
	}

	// 使用默认键函数
	return DefaultCacheKeyFunc(options.Method, fullURL, options.Headers)
}

// getCacheTTL 确定缓存响应的 TTL
func getCacheTTL(cacheConfig *CacheConfig, options *AxiosOptions) time.Duration {
	// 每次请求的 TTL 优先
	if options.Cache != nil && options.Cache.TTL > 0 {
		return options.Cache.TTL
	}

	// 回退到客户端默认 TTL
	return cacheConfig.DefaultTTL
}

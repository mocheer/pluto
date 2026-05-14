package axios

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// LogLevel 表示日志级别
type LogLevel int

const (
	// LevelNone 表示不记录日志
	LevelNone LogLevel = iota
	// LevelError 表示仅记录错误日志
	LevelError
	// LevelInfo 表示记录信息级别日志
	LevelInfo
	// LevelDebug 表示记录调试级别日志
	LevelDebug
)

// Logger 是日志记录器接口
type Logger interface {
	LogRequest(*http.Request)
	LogResponse(*http.Response, []byte, time.Duration)
	LogError(error)
	SetLevel(LogLevel)
}

// LogOptions 包含日志记录器的配置选项
type LogOptions struct {
	Level          LogLevel
	MaxBodyLength  int
	MaskHeaders    []string
	Output         io.Writer
	TimeFormat     string
	IncludeBody    bool
	IncludeHeaders bool
}

// DefaultLogger 是默认的日志记录器实现
type DefaultLogger struct {
	options LogOptions
}

// NewDefaultLogger 创建一个新的默认日志记录器
func NewDefaultLogger(options LogOptions) *DefaultLogger {
	if options.Output == nil {
		options.Output = os.Stdout
	}
	if options.TimeFormat == "" {
		options.TimeFormat = time.RFC3339
	}
	if options.MaxBodyLength == 0 {
		options.MaxBodyLength = 1000
	}
	return &DefaultLogger{options: options}
}

// SetLevel 设置日志级别
func (l *DefaultLogger) SetLevel(level LogLevel) {
	l.options.Level = level
}

// LogRequest 记录 HTTP 请求日志
func (l *DefaultLogger) LogRequest(req *http.Request) {
	if l.options.Level > LevelNone {
		return
	}

	var buf strings.Builder
	timestamp := time.Now().Format(l.options.TimeFormat)

	fmt.Fprintf(&buf, "[%s] REQUEST: %s %s\n", timestamp, req.Method, req.URL)

	if l.options.IncludeHeaders {
		buf.WriteString("Headers:\n")
		for key, vals := range req.Header {
			if l.isHeaderMasked(key) {
				fmt.Fprintf(&buf, "  %s: [MASKED]\n", key)
			} else {
				fmt.Fprintf(&buf, "  %s: %s\n", key, strings.Join(vals, ", "))
			}
		}
	}

	if l.options.IncludeBody && req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err == nil {
			req.Body = io.NopCloser(bytes.NewBuffer(body))
			if len(body) > l.options.MaxBodyLength {
				fmt.Fprintf(&buf, "Body: (truncated) %s...\n", body[:l.options.MaxBodyLength])
			} else {
				fmt.Fprintf(&buf, "Body: %s\n", body)
			}
		}
	}

	fmt.Fprintln(l.options.Output, buf.String())
}

// LogResponse 记录 HTTP 响应日志
func (l *DefaultLogger) LogResponse(resp *http.Response, body []byte, duration time.Duration) {
	if l.options.Level > LevelNone {
		return
	}

	var buf strings.Builder
	timestamp := time.Now().Format(l.options.TimeFormat)

	fmt.Fprintf(&buf, "[%s] RESPONSE: %d %s (%.2fms)\n",
		timestamp, resp.StatusCode, resp.Status, float64(duration.Microseconds())/1000)

	if l.options.IncludeHeaders {
		buf.WriteString("Headers:\n")
		for key, vals := range resp.Header {
			if l.isHeaderMasked(key) {
				fmt.Fprintf(&buf, "  %s: [MASKED]\n", key)
			} else {
				fmt.Fprintf(&buf, "  %s: %s\n", key, strings.Join(vals, ", "))
			}
		}
	}

	if l.options.IncludeBody && body != nil {
		if len(body) > l.options.MaxBodyLength {
			fmt.Fprintf(&buf, "Body: (truncated) %s...\n", body[:l.options.MaxBodyLength])
		} else {
			fmt.Fprintf(&buf, "Body: %s\n", body)
		}
	}

	fmt.Fprintln(l.options.Output, buf.String())
}

// LogError 记录错误日志
func (l *DefaultLogger) LogError(err error) {
	if l.options.Level > LevelError {
		return
	}

	timestamp := time.Now().Format(l.options.TimeFormat)
	fmt.Fprintf(l.options.Output, "[%s] ERROR: %v\n", timestamp, err)
}

// isHeaderMasked 检查请求头是否需要被脱敏
func (l *DefaultLogger) isHeaderMasked(header string) bool {
	header = strings.ToLower(header)
	for _, masked := range l.options.MaskHeaders {
		if strings.ToLower(masked) == header {
			return true
		}
	}
	return false
}

// NewLogger 创建一个新的日志记录器，使用默认配置
func NewLogger(level LogLevel) Logger {
	return NewDefaultLogger(LogOptions{
		Level:          level,
		MaxBodyLength:  1000,
		MaskHeaders:    []string{"Authorization", "Cookie", "Set-Cookie"},
		Output:         os.Stdout,
		TimeFormat:     time.RFC3339,
		IncludeBody:    true,
		IncludeHeaders: true,
	})
}

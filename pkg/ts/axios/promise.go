package axios

import "sync"

// Promise 表示异步 HTTP 请求的 Promise
type Promise struct {
	response *Response
	err      error
	then     func(*Response)
	catch    func(error)
	finally  func()
	done     chan struct{}
	mu       sync.Mutex
}

// Then 注册 Promise 成功时的回调函数
func (p *Promise) Then(fn func(*Response)) *Promise {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.response != nil && p.err == nil {
		fn(p.response)
	} else {
		p.then = fn
	}
	return p
}

// Catch 注册 Promise 失败时的回调函数
func (p *Promise) Catch(fn func(error)) *Promise {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.err != nil {
		fn(p.err)
	} else {
		p.catch = fn
	}
	return p
}

// Finally 注册 Promise 完成时的回调函数（无论成功或失败）
func (p *Promise) Finally(fn func()) {
	p.mu.Lock()

	if p.response != nil || p.err != nil {
		p.mu.Unlock()
		fn()
	} else {
		p.finally = fn
		p.mu.Unlock()
	}

	<-p.done
}

// NewPromise 创建一个新的 Promise
func NewPromise() *Promise {
	return &Promise{
		done: make(chan struct{}),
	}
}

// resolve 解析 Promise，设置响应或错误并触发回调
func (p *Promise) resolve(resp *Response, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.response = resp
	p.err = err

	if p.then != nil && err == nil {
		p.then(resp)
	}
	if p.catch != nil && err != nil {
		p.catch(err)
	}
	if p.finally != nil {
		p.finally()
	}

	close(p.done)
}

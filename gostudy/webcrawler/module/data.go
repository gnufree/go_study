package module

import "net/http"

// 数据请求的类型
type Request struct {
	// HTTP请求
	httpReq *http.Request
	// 请求的深度
	depth uint32
}

// 数据响应的类型
type Response struct {
	// HTTP 响应
	httpResp *http.Response
	// 响应的深度
	depth uint32
}
// 用于创建一个新的请求实例
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq:httpReq,depth:depth}
}

// 用户获取HTTP 请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// 用户获取请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}


package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewResponse 返回响应
func NewResponse(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// NewSuccessResponse 返回成功响应
func NewSuccessResponse(data interface{}, c *gin.Context) {
	NewResponse(200, "success", data, c)
}

// NewFailResponse 返回失败响应
func NewFailResponse(msg string, c *gin.Context) {
	NewResponse(400, msg, nil, c)
}

// NewUnauthorizedResponse 返回未授权响应
func NewUnauthorizedResponse(c *gin.Context) {
	NewResponse(401, "unauthorized", nil, c)
}

// NewForbiddenResponse 返回禁止响应
func NewForbiddenResponse(c *gin.Context) {
	NewResponse(403, "forbidden", nil, c)
}

// NewNotFoundResponse 返回未找到响应
func NewNotFoundResponse(c *gin.Context) {
	NewResponse(404, "not found", nil, c)
}

// NewMethodNotAllowedResponse 返回方法不允许响应
func NewMethodNotAllowedResponse(c *gin.Context) {
	NewResponse(405, "method not allowed", nil, c)
}

// NewNotAcceptableResponse 返回不可接受响应
func NewNotAcceptableResponse(c *gin.Context) {
	NewResponse(406, "not acceptable", nil, c)
}

// NewConflictResponse 返回冲突响应
func NewConflictResponse(c *gin.Context) {
	NewResponse(409, "conflict", nil, c)
}

// NewUnsupportedMediaTypeResponse 返回不支持的媒体类型响应
func NewUnsupportedMediaTypeResponse(c *gin.Context) {
	NewResponse(415, "unsupported media type", nil, c)
}

// NewTooManyRequestsResponse 返回请求过多响应
func NewTooManyRequestsResponse(c *gin.Context) {
	NewResponse(429, "too many requests", nil, c)
}

// NewInternalServerErrorResponse 返回服务器内部错误响应
func NewInternalServerErrorResponse(c *gin.Context) {
	NewResponse(500, "internal server error", nil, c)
}

// NewServiceUnavailableResponse 返回服务不可用响应
func NewServiceUnavailableResponse(c *gin.Context) {
	NewResponse(503, "service unavailable", nil, c)
}

// NewGatewayTimeoutResponse 返回网关超时响应
func NewGatewayTimeoutResponse(c *gin.Context) {
	NewResponse(504, "gateway timeout", nil, c)
}

// NewCodeResponse 返回自定义Code响应(验证码)
func NewCodeResponse(code string, msg string, c *gin.Context) {
	NewResponse(200, msg, code, c)
}

// NewRedirectResponse 返回重定向响应
func NewRedirectResponse(code int, location string, c *gin.Context) {
	c.Redirect(code, location)
}

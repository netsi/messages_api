package test

import (
	"bufio"
	"net"
	"net/http"
)

type ResponseWriterMock struct {
	HTTPHeader   http.Header
	StatusCode   int
	ResponseBody []byte
}

func NewResponseWritterMock() *ResponseWriterMock {
	return &ResponseWriterMock{
		HTTPHeader: http.Header{},
	}
}

func (r *ResponseWriterMock) Header() http.Header {
	return r.HTTPHeader
}

func (r *ResponseWriterMock) Write(bytes []byte) (int, error) {
	r.ResponseBody = bytes
	return len(bytes), nil
}

func (r *ResponseWriterMock) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
}

func (r *ResponseWriterMock) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (r *ResponseWriterMock) Flush() {
}

func (r *ResponseWriterMock) CloseNotify() <-chan bool {
	return nil
}

func (r *ResponseWriterMock) Status() int {
	return r.StatusCode
}

func (r *ResponseWriterMock) Size() int {
	return 0
}

func (r *ResponseWriterMock) WriteString(s string) (int, error) {
	r.ResponseBody = []byte(s)
	return len(r.ResponseBody), nil
}

func (r *ResponseWriterMock) Written() bool {
	return len(r.ResponseBody) > 0
}

func (r *ResponseWriterMock) WriteHeaderNow() {

}

func (r *ResponseWriterMock) Pusher() http.Pusher {
	return nil
}

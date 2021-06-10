package mocks

import (
	"bufio"
	"net"
	"net/http"

	"github.com/stretchr/testify/mock"
)

// ResponseWriterMock facilitates testability of gin contexts and controllers.
type ResponseWriterMock struct {
	mock.Mock
}

// Status returns the HTTP response status code of the current request.
func (w *ResponseWriterMock) Status() int {
	args := w.Called()
	return args.Get(0).(int)
}

// Size returns the number of bytes already written into the response http body.
// See Written()
func (w *ResponseWriterMock) Size() int {
	args := w.Called()
	return args.Get(0).(int)
}

// WriteString writes the string into the response body.
func (w *ResponseWriterMock) WriteString(str string) (int, error) {
	args := w.Called(str)

	numBytes := args.Get(0).(int)
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return numBytes, err
}

// Written returns true if the response body was already written.
func (w *ResponseWriterMock) Written() bool {
	args := w.Called()
	result := false
	if args.Get(0) != nil {
		result = args.Get(0).(bool)
	}

	return result
}

// WriteHeaderNow forces to write the http header (status code + headers).
func (w *ResponseWriterMock) WriteHeaderNow() {
	w.Called()
}

// Pusher get the http.Pusher for server push
func (w *ResponseWriterMock) Pusher() http.Pusher {
	panic("not implemented yet ...")
}

// CloseNotify closes the underlining communication channel.
func (w *ResponseWriterMock) CloseNotify() <-chan bool {
	panic("not implemented yet ...")
}

// Flush sends any buffered data to the client.
func (w *ResponseWriterMock) Flush() {
	w.Called()
}

// Hijack lets the caller take over the connection.
func (w *ResponseWriterMock) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	panic("not implemented yet ...")
}

// Header returns the header map that will be sent by
// WriteHeader. The Header map also is the mechanism with which
// Handlers can set HTTP trailers.
func (w *ResponseWriterMock) Header() http.Header {
	args := w.Called()
	var headers map[string][]string = nil
	if args.Get(0) != nil {
		headers = args.Get(0).(map[string][]string)
	}

	return headers
}

// Write writes the data to the connection as part of an HTTP reply.
func (w *ResponseWriterMock) Write(buf []byte) (int, error) {
	args := w.Called(buf)

	var numBytes int
	if args.Get(0) != nil {
		numBytes = args.Get(0).(int)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return numBytes, err
}

// WriteHeader sends an HTTP response header with the provided
// status code.
func (w *ResponseWriterMock) WriteHeader(statusCode int) {
	w.Called(statusCode)
}

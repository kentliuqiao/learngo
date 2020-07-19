package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func errPanic(w http.ResponseWriter,
	r *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter,
	r *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(w http.ResponseWriter,
	r *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(w http.ResponseWriter,
	r *http.Request) error {
	return os.ErrPermission
}

func errUnknown(w http.ResponseWriter,
	r *http.Request) error {
	return errors.New("nnknown error")
}

func noError(w http.ResponseWriter,
	r *http.Request) error {
	fmt.Fprintf(w, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error\n"},
	{errUserError, 400, "user error\n"},
	{errNotFound, 404, "Not Found\n"},
	{errNoPermission, 403, "Forbidden\n"},
	{errUnknown, 500, "Internal Server Error\n"},
	{noError, 200, "no error"},
}

// 测试errWrapper函数
func TestErrorWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet, "http://imooc.com", nil)

		f(resp, req)
		verifyResponse(resp.Result(), tt.code, tt.message, t)
	}
}

// 测试服务器
func TestErrorWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(
			http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != expectedCode ||
		string(body) != expectedMsg {
		t.Errorf("expect (%d, %s); got(%d, %s)",
			expectedCode, expectedMsg, resp.StatusCode, body)
	}
}

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}


// 公用测试数据
var tests = []struct {
	// 期望的输入是一个appHandler
	h appHandler
	// 期望的输出，包含一个code和一个msg
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 测试errWrapper方法的正确性
func TestErrWrapper(t *testing.T)  {
	for _, tt := range tests {
		// 目标函数为f
		f := errWrapper(tt.h)
		// 测试目标函数的行为
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 启动一个服务器并在服务器中测试
// 和之前的不同：之前是直接调用方法，而没有服务器
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	// 原本的body中最后有一个换行符，要把它trim掉
	body := strings.Trim(string(b), "\n")

	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); " + "got (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
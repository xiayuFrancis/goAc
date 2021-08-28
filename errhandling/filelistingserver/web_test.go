package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}


func Test_errWrapper(t *testing.T) {
	type args struct {
		handler appHandler
	}
	tests := []struct {
		name    string
		args    args
		code    int
		message string
	}{
		{"errpanic", args{errPanic}, 500, "Internal Server Error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := errWrapper(tt.args.handler)
			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
			f(response, request)
			body, _ := ioutil.ReadAll(response.Body)
			if response.Code != tt.code || strings.Trim(string(body), "\n") != tt.message {
				t.Errorf("want (%d, %s), got (%d, %s)", tt.code, tt.message, response.Code, string(body))
			}
		})
	}
}

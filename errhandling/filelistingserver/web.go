package main

import (
	"goAc/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
	"runtime"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				log.Printf("http: panic serving %v: %v\n%s", request.RemoteAddr, err, buf)
			}
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()

		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			log.Println(err)
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

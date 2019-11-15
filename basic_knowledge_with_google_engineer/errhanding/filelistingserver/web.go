package main

import (
	"log"
	"my_go_demo/learn_go2/errhanding/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 函数式编程，输入是一个函数，输出是一个函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		// user error
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			// system error
			code := http.StatusOK
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

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.FileHandler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

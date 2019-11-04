package main

import (
	"LearnGo/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

// 错误处理结构：
// appHandler是方法的别名，是errWrapper()的参数，其结构和HandleFileList()相同
// handler.go中定义的HandleFileList()执行业务代码，碰到err就return
// errWrapper()的参数和返回值都是一个函数，
// 它包装HandleFileList()，在返回值中执行appHandler，并进行相应的错误处理
// errWrapper()的返回值作为http.HandleFunc()的第二个参数

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil { // 必须要在这里判断r是否为nil，否则即使正常运行，下面的语句也会被执行到
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf("Panic: %v", r)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): // 文件不存在错误
				code = http.StatusNotFound
			case os.IsPermission(err): // 权限错误
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			// 第一个参数writer表示向谁汇报错误
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

// 可以暴露给用户看的error
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

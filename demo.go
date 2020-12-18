package Unified_exception_handling_demo

import (
	"log"
	"net/http"
	"os"
)

//声明一个函数类型
//作为一个占位用
type appHandel func(writer http.ResponseWriter, request *http.Request) error

//函数式编程
func errWrapper (handel appHandel) func(writer http.ResponseWriter, request *http.Request){
	
	return func(writer http.ResponseWriter, request *http.Request) {
		//异常捕获
		defer func() {
			if r:= recover();r!=nil {
				log.Panicf("Panic: %v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
	}
	
	err :=handel(writer, request)
	if err != nil {
		
		//业务异常展示给用户看
		if uerError, ok := err.(UserError); ok{
			http.Error(writer,uerError.Message(),http.StatusBadRequest)
			return nil
		}

		code := http.StatusOK
		switch  {
			case os.IsNotExist(err):
				code = http.StatusInternalServerError
			default:
				code = http.StatusInternalServerError
		}
		http.Error(writer,"Sorry,FileNotFound!",code)
	}
}
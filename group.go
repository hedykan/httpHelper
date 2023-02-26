package httpHelper

import (
	"net/http"
	"strings"
)

type Handle struct {
	Url     string
	Handler http.Handler
}

type HandleArr []Handle

// 新增url组
func (arr HandleArr) AddGroup(parentUrl string) HandleArr {
	for i := 0; i < len(arr); i++ {
		parentUrl = formatUrl(parentUrl)
		arr[i].Url = formatUrl(arr[i].Url)
		arr[i].Url = parentUrl + "/" + arr[i].Url
	}
	return arr
}

// 新增中间件
func (arr HandleArr) AddMiddleward(middleward Middleware, param ...interface{}) HandleArr {
	for i := 0; i < len(arr); i++ {
		arr[i].Handler = middleward(arr[i].Handler, param...)
	}
	return arr
}

// 设置请求方法
func (arr HandleArr) SetMethod(method string) HandleArr {
	for i := 0; i < len(arr); i++ {
		arr[i].Handler = methodMiddleware(arr[i].Handler, method)
	}
	return arr
}

// 设置请求复用器
func SetMuxHandle(mux *http.ServeMux, handleArr HandleArr) {
	SetMuxHandleAddMiddleware(mux, handleArr, CrosMiddleward)
}

// 设置请求复用器并添加中间件
func SetMuxHandleAddMiddleware(mux *http.ServeMux, handleArr HandleArr, middlewareArr ...Middleware) {
	for _, v := range middlewareArr {
		handleArr.AddMiddleward(v)
	}
	for i := 0; i < len(handleArr); i++ {
		if handleArr[i].Url[0] != '/' {
			handleArr[i].Url = "/" + handleArr[i].Url
		}
		mux.Handle(handleArr[i].Url, handleArr[i].Handler)
	}
}

func formatUrl(str string) string {
	strArr := strings.Split(str, "/")
	str = ""
	for i := 0; i < len(strArr); i++ {
		if strArr[i] != "" {
			if str == "" {
				str = strArr[i]
			} else {
				str = str + "/" + strArr[i]
			}
		}
	}
	return str
}

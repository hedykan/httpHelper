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

func (arr HandleArr) AddGroup(parentUrl string) HandleArr {
	for i := 0; i < len(arr); i++ {
		parentUrl = formatUrl(parentUrl)
		arr[i].Url = formatUrl(arr[i].Url)
		arr[i].Url = parentUrl + "/" + arr[i].Url
	}
	return arr
}

func (arr HandleArr) AddMiddleward(middleward Middleware, param ...interface{}) HandleArr {
	for i := 0; i < len(arr); i++ {
		arr[i].Handler = middleward(arr[i].Handler, param...)
	}
	return arr
}

func (arr HandleArr) SetMethod(method string) HandleArr {
	for i := 0; i < len(arr); i++ {
		arr[i].Handler = methodMiddleware(arr[i].Handler, method)
	}
	return arr
}

func SetMuxHandle(mux *http.ServeMux, handleArr HandleArr) {
	handleArr.AddMiddleward(crosMiddleward)
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

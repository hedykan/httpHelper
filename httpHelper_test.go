package httpHelper

import (
	"fmt"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	mux := http.NewServeMux()
	handleArr := HandleArr{
		{
			Url:     "index",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { Write(w, "hello world, im in index") }),
		},
		{
			Url:     "page",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { Write(w, "hello world, im in page") }),
		},
	}
	handleArr.AddGroup("get").SetMethod(http.MethodGet)

	SetMuxHandle(mux, handleArr)
	fmt.Println(mux)
}

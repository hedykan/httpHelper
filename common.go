package httpHelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/thinkeridea/go-extend/exnet"
)

func RemoteIp(req *http.Request) string {
	ip := exnet.ClientPublicIP(req)
	if ip == "" {
		ip = exnet.ClientIP(req)
	}
	return ip
}

func Write(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(data)
	enc.Encode(res)
}

func WriteList(w http.ResponseWriter, list interface{}, count int) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := resList(list, count)
	enc.Encode(res)
}

func WriteError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.Encode(map[string]interface{}{"code": code, "msg": msg})
}

func Get(r *http.Request) map[string]string {
	var res = make(map[string]string)
	keys := r.URL.Query()
	for index, value := range keys {
		res[index] = value[0]
	}

	return res
}

func PostJson(r *http.Request, obj interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// 重新写入
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, obj)
	// err := json.NewDecoder(r.Body).Decode(obj) // 会导致r.Body读取完后无法重新写入
	if err != nil {
		panic(err)
	}
}

func res(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	if data != nil {
		res["data"] = data
	}
	res["msg"] = "ok"
	return res
}

func resList(list interface{}, count int) map[string]interface{} {
	data := make(map[string]interface{})
	data["list"] = list
	data["count"] = count
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["data"] = data
	res["msg"] = "ok"
	return res
}

func GetPageQuery(r *http.Request) (int, int, error) {
	query := Get(r)
	page, err := strconv.Atoi(query["page"])
	if err != nil {
		return 0, 0, err
	}
	size, err := strconv.Atoi(query["page"])
	if err != nil {
		return 0, 0, err
	}
	return page, size, nil
}

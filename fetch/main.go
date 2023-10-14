package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	// 定義資料類型
	type ReqParam struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}

	type ResParam struct {
		D string `json:"d"`
		E string `json:"e"`
		F string `json:"f"`
	}

	type DBParam struct {
		A string
		B string
	}

	// 變數初始化
	reqParam := &ReqParam{}
	resParam := &ResParam{}
	dbParam := &DBParam{}
	otherParam := &DBParam{}

	// 程式開始
	url := "http://localhost:5678"
	payload := `{"a":"aaa", "b":"bbb","c":"ccc"}`

	fetch(url, payload).
		printJSON().
		bind(&reqParam).
		call(loginLogic, reqParam).
		call(loginDB, dbParam).
		call(loginOther, otherParam).
		respond(resParam)

}

func loginLogic(v any) {
	fmt.Println("login logic")
}

func loginDB(v any) {
	fmt.Println("login database")
}

func loginOther(v any) {
	fmt.Println("login other")
}

type req struct {
	resp *http.Response
	body []byte
}

func fetch(url string, body string) *req {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Println("http post failed")
		return nil
	}

	r := new(req)
	r.resp = resp

	return r
}

func (r *req) respond(v any) *req {
	fmt.Println("respond to target")
	return r
}

func (r *req) print() *req {
	body, err := io.ReadAll(r.resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	r.body = body

	fmt.Println(string(r.body))

	return r
}

func (r *req) printJSON() *req {
	data := make(map[string]string)

	err := json.Unmarshal(r.body, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(json))

	return r
}

func (r *req) bind(v any) *req {

	err := json.Unmarshal(r.body, v)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return r
}

func (r *req) call(fn func(v any), v any) *req {
	fn(v)
	return r
}

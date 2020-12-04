package camp

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// send_request

type RequestAPI interface {
	Url() string
	Method() string
	ContentType() string
	// 请求参数
	RequestParams() (io.Reader, error)
	// 处理响应
	// 返回处理结果 也可传入指针接收数据
	// string指示响应body的content-type (用于响应type不唯一的情况)
	ParseResponse([]byte, string, interface{}) (interface{}, error)
}

func SendRequest(rd RequestAPI, model interface{}) (interface{}, error) {
	params, err := rd.RequestParams()
	if err != nil {
		return nil, err
	}
	resp, contentType, err := newRequest(rd.Method(), rd.Url(), rd.ContentType(), params)
	if err != nil {
		return nil, err
	}
	return rd.ParseResponse(resp, contentType, model)
}

// method请求方法 url路径 contentType请求类型
func newRequest(method, url, bodyType string, body io.Reader) ([]byte, string, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, "", err
	}
	if bodyType != "" {
		req.Header.Set("content-type", bodyType)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	// 响应类型
	var typ string
	dataType, ok := resp.Header["Content-Type"]
	if ok && len(dataType) > 0 {
		typ = dataType[0]
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err1 := resp.Body.Close(); err == nil {
		err = err1
	}
	return data, typ, err
}

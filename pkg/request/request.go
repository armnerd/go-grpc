package request

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/idoubi/goz"
)

// Get 请求
func Get(url string, data interface{}, headers map[string]interface{}) (string, error) {
	var res string
	cli := goz.NewClient()
	resp, err := cli.Get(url, goz.Options{
		Headers: headers,
		Query:   data,
	})
	if err != nil {
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		return res, err
	}
	return body.GetContents(), nil
}

// PostForm 请求
func PostForm(url string, data map[string]interface{}, headers map[string]interface{}) (string, error) {
	var res string
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Headers:    headers,
		FormParams: data,
	})
	if err != nil {
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		return res, err
	}
	return body.GetContents(), nil
}

// PostJson 请求
func PostJson(url string, data string, headers map[string]string) (string, error) {
	var res string
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	for index := range headers {
		req.Header.Set(index, headers[index])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(content), nil
}

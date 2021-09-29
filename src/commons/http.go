package commons

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpUtil struct {
}


// PostForm 封装PostForm请求
func (c *HttpUtil) PostForm(url string, data url.Values) (result string, err error) {
	h := http.Client{}
	resp, err := h.PostForm(url, data)
	if err != nil{
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), err
}



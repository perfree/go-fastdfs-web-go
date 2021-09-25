package commons

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpUtil struct {
}

// Get 封装Get请求
func (c *HttpUtil) Get(url string) string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

// PostForm 封装PostForm请求
func (c *HttpUtil) PostForm(url string, data url.Values) (result string, err error) {
	h := http.Client{}
	resp, err := h.PostForm(url, data)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), err
}

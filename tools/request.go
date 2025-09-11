package tools

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// HTTPRequest 封装的 HTTP 请求方法
func HTTPRequest(method, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	// 处理 body 为空的情况
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}

	// 创建请求对象
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, 0, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建客户端
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, nil
}

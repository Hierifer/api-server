package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"encoding/json"
)

func Request(url string, body map[string]interface{}, optionalHeaders ...map[string]string) (string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
    // 添加可选的 headers
    if len(optionalHeaders) > 0 {
        for key, value := range optionalHeaders[0] {
            req.Header.Set(key, value)
        }
    }
    // 发送请求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
	defer resp.Body.Close() // 确保响应体被关闭
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(respBody), nil
}
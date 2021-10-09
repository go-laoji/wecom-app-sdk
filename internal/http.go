package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BizResponse 基础返回类型，定义错误代码及错误消息
type BizResponse struct {
	Error
}

const (
	qyApiHost = "https://qyapi.weixin.qq.com"
)

func httpRequestWithContext(ctx context.Context, request *http.Request, resChan chan<- []byte) (err error) {
	request = request.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("client.Do Error: %s", err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll Error: %s", err.Error())
	}
	defer resp.Body.Close()
	resChan <- data
	return nil
}

func HttpGet(apiUrl string) (body []byte, err error) {
	resChan := make(chan []byte)
	repoUrl := fmt.Sprintf("%s%s", qyApiHost, apiUrl)
	request, err := http.NewRequest(http.MethodGet, repoUrl, nil)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(15)*time.Second))
	defer cancel()

	go httpRequestWithContext(ctx, request, resChan)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.Tick(time.Duration(15) * time.Second):
		return nil, errors.New("time over")
	case body = <-resChan:
		return body, nil
	}
}

func HttpPost(apiUrl string, params interface{}) (body []byte, err error) {
	resChan := make(chan []byte)
	repoUrl := fmt.Sprintf("%s%s", qyApiHost, apiUrl)
	data, err := json.Marshal(params)
	//log.Println(string(data))
	request, err := http.NewRequest(http.MethodPost, repoUrl, bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(15)*time.Second))
	defer cancel()

	go httpRequestWithContext(ctx, request, resChan)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.Tick(time.Duration(15) * time.Second):
		return nil, errors.New("time over")
	case body = <-resChan:
		return body, nil
	}
}

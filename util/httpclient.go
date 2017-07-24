package util

import (
	"encoding/json"

	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/yl10/teambition/model"
)

// HttpClient 混入, 可直接调用其方法
type HttpClient struct {
	Token string
	*gorequest.SuperAgent
}

// NewHttpClient 新建请求客户端
func NewHttpClient() *HttpClient {
	cli := gorequest.New()
	return &HttpClient{SuperAgent: cli}
}

func (t *HttpClient) SetToken(token string) *HttpClient {
	t.Token = token
	return t
}

// Get 请求
func (t *HttpClient) Get(url string, params interface{}) ([]byte, error) {
	res, dt, err := t.SuperAgent.Get(url).Send(params).Set("Authorization", "OAuth2 "+t.Token).EndBytes()
	if err != nil {
		return nil, TransErrs2XErr(err)
	}

	return t.checkResult(res, dt)
}

// GetReq Get请求，返回 response
func (t *HttpClient) GetReq(url string, params interface{}) (*http.Response, error) {
	res, _, err := t.SuperAgent.Get(url).Send(params).Set("Authorization", "OAuth2 "+t.Token).EndBytes()
	if err != nil {
		return nil, JoinErrs(err)
	}

	return res, nil
}

// Post 请求
func (t *HttpClient) Post(url string, params interface{}) ([]byte, error) {
	res, dt, err := t.SuperAgent.Post(url).SendStruct(params).Set("Authorization", "OAuth2 "+t.Token).EndBytes()
	if err != nil {
		return nil, TransErrs2XErr(err)
	}

	return t.checkResult(res, dt)
}

// Put 请求
func (t *HttpClient) Put(url string, params interface{}) ([]byte, error) {
	res, dt, err := t.SuperAgent.Put(url).SendStruct(params).Set("Authorization", "OAuth2 "+t.Token).EndBytes()
	if err != nil {
		return nil, TransErrs2XErr(err)
	}

	return t.checkResult(res, dt)
}

// Delete 请求
func (t *HttpClient) Delete(url string, params interface{}) ([]byte, error) {
	res, dt, err := t.SuperAgent.Delete(url).Send(params).Set("Authorization", "OAuth2 "+t.Token).EndBytes()
	if err != nil {
		return nil, TransErrs2XErr(err)
	}

	return t.checkResult(res, dt)
}

func (t *HttpClient) checkResult(res gorequest.Response, dt []byte) ([]byte, error) {
	if res.StatusCode >= http.StatusMultipleChoices || res.StatusCode < http.StatusOK {
		var tbErr model.TBAPIError
		if err := json.Unmarshal(dt, &tbErr); err != nil {
			return nil, NewError(string(dt))
		}

		return nil, NewError(tbErr.Message, tbErr.Code)
	}

	return dt, nil
}

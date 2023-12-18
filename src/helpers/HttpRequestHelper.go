package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"io/ioutil"
	"net/http"
)

type httpRequestHelper struct {
	client *http.Client
	url    string
}

func NewHttpRequestHelper(ctx context.Context, url string) (*httpRequestHelper, *models.MyError) {
	return &httpRequestHelper{
		url:    url,
		client: &http.Client{},
	}, nil
}

func (rh *httpRequestHelper) send(method string, endpoint string, params map[string]string, body interface{}) *models.Response {
	if params == nil {
		params = map[string]string{}
	}

	var jsonValue []byte
	if body == nil {
		jsonValue = []byte{}
	} else {
		jsonValue, _ = json.Marshal(body)
	}
	url := rh.url + endpoint

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return &models.Response{
			Code:    0,
			Content: nil,
			Error:   []*models.MyError{response.GetError(err)},
		}
	}
	q := req.URL.Query()
	for s, s2 := range params {
		q.Add(s, s2)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := rh.client.Do(req)
	if err != nil {
		return &models.Response{
			Code:    0,
			Content: nil,
			Error:   []*models.MyError{response.GetError(err)},
		}
	}
	defer res.Body.Close()
	by, err := ioutil.ReadAll(res.Body)
	return &models.Response{
		Code:    res.StatusCode,
		Content: string(by),
		Error:   []*models.MyError{},
	}
}

func (rh *httpRequestHelper) Post(endpoint string, params map[string]string, body interface{}) *models.Response {
	return rh.send("POST", endpoint, params, body)
}

func (rh *httpRequestHelper) Get(endpoint string, params map[string]string, body interface{}) *models.Response {
	return rh.send("GET", endpoint, params, body)
}

func (rh *httpRequestHelper) Put(endpoint string, params map[string]string, body interface{}) *models.Response {
	return rh.send("PUT", endpoint, params, body)
}

func (rh *httpRequestHelper) Delete(endpoint string, params map[string]string, body interface{}) *models.Response {
	return rh.send("DELETE", endpoint, params, body)
}

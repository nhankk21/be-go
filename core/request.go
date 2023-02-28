package core

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	request *http.Request
}

func (request *Request) GetParams() map[string]string {
	params := make(map[string]string)
	for k := range request.request.URL.Query() {
		params[k] = request.request.URL.Query().Get(k)
	}
	return params
}
func (request *Request) GetParam(k string) string {
	return request.request.URL.Query().Get(k)
}
func (request *Request) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for k := range request.request.Header {
		headers[k] = request.request.Header.Get(k)
	}
	return headers
}
func (request *Request) GetHeader(k string) string {
	return request.request.Header.Get(k)
}

func (request *Request) GetBody(body interface{}) error {
	// var data interface{}
	if err := json.NewDecoder(request.request.Body).Decode(&body); err != nil {
		return err
	}
	return nil
	// err := json.NewDecoder(request.request.Body).Decode(&data)
	// if err != nil {
	// 	return err
	// }
	// jsonStr, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }
	// err = json.Unmarshal(jsonStr, &body)
	// return err
}
func (request *Request) GetBodyMap() (map[string]interface{}, error) {
	// var data interface{}
	// err := json.NewDecoder(request.request.Body).Decode(&data)
	// if err != nil {
	// 	return nil, err
	// }
	// jsonStr, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }
	m := make(map[string]interface{})
	if err := json.NewDecoder(request.request.Body).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
	// err = json.Unmarshal(jsonStr, &m)
	// if err != nil {
	// 	return nil, err
	// }
	// return m, nil
}

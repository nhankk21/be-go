package common

import (
	"encoding/json"
	"io"
)

func GetBodyContent(req io.ReadCloser, body interface{}) error {

	if err := json.NewDecoder(req).Decode(&body); err != nil {
		return err
	}
	return nil
	// var data interface{}
	// err := json.NewDecoder(req).Decode(&data)
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

func GetBodyContentMap(req io.ReadCloser) (map[string]interface{}, error) {

	m := make(map[string]interface{})
	if err := json.NewDecoder(req).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
	// var data interface{}
	// err := json.NewDecoder(req).Decode(&data)
	// if err != nil {
	// 	return nil, err
	// }
	// jsonStr, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }
	// m := make(map[string]interface{})
	// err = json.Unmarshal(jsonStr, &m)
	// if err != nil {
	// 	return nil, err
	// }
	// return m, nil
}

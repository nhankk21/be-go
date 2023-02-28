package core

type Response struct {
	Status  APIStatus   `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

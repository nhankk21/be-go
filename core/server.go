package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Server struct{}

func NewServer() *Server {
	hostname, err := os.Hostname()
	if err != nil {
		return nil
	}
	fmt.Println("Host name: " + hostname)
	return &Server{}
}

func (server *Server) Expose(port string) error {
	fmt.Println("Port: " + port)
	err := http.ListenAndServe(port, nil)
	return err
}

// func (server *Server) SetHandler(path string, f func(Request) *Response, _method APIMethod) {
// 	var method string
// 	switch _method {
// 	case APIMethodEnum.GET:
// 		method = http.MethodGet
// 	case APIMethodEnum.POST:
// 		method = http.MethodPost
// 	case APIMethodEnum.PUT:
// 		method = http.MethodPut
// 	case APIMethodEnum.DELETE:
// 		method = http.MethodDelete
// 	}
// 	fu := func(w http.ResponseWriter, req *http.Request) {
// 		request := Request{
// 			request: req,
// 		}
// 		reps := f(request)
// 		var status int
// 		switch reps.Status {
// 		case APIStatusEnum.Ok:
// 			status = http.StatusOK
// 		case APIStatusEnum.Error:
// 			status = http.StatusBadRequest
// 		case APIStatusEnum.Invalid:
// 			status = http.StatusBadRequest
// 		case APIStatusEnum.NotFound:
// 			status = http.StatusNotFound
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(status)
// 		json.NewEncoder(w).Encode(reps)
// 		// responseWithJson(w, status, reps)
// 	}
// 	// server.router.HandleFunc(path, fu).Methods(method)
// }

func (server *Server) SetHandler(path string, f func(Request) *Response, _method APIMethod) {
	var method string
	switch _method {
	case APIMethodEnum.GET:
		method = http.MethodGet
	case APIMethodEnum.POST:
		method = http.MethodPost
	case APIMethodEnum.PUT:
		method = http.MethodPut
	case APIMethodEnum.DELETE:
		method = http.MethodDelete
	}
	fu := func(w http.ResponseWriter, req *http.Request) {
		_method := req.Method
		if _method != method {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Method not allowed",
			})
			return
		}
		request := Request{
			request: req,
		}
		reps := f(request)
		var status int
		switch reps.Status {
		case APIStatusEnum.Ok:
			status = http.StatusOK
		case APIStatusEnum.Error:
			status = http.StatusBadRequest
		case APIStatusEnum.Invalid:
			status = http.StatusBadRequest
		case APIStatusEnum.NotFound:
			status = http.StatusNotFound
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(reps)
		// responseWithJson(w, status, reps)
	}
	// server.router.HandleFunc(path, fu).Methods(method)
	http.HandleFunc(path, fu)
}

// func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	writer.WriteHeader(status)
// 	json.NewEncoder(writer).Encode(object)
// }

package core

type APIMethod string

type apiMethod struct {
	GET     APIMethod
	POST    APIMethod
	PUT     APIMethod
	DELETE  APIMethod
	OPTIONS APIMethod
}

var APIMethodEnum = &apiMethod{
	GET:    "GET",
	POST:   "POST",
	PUT:    "PUT",
	DELETE: "DELETE",
}

type APIStatus string

type apiStatus struct {
	Ok           APIStatus
	Error        APIStatus
	Invalid      APIStatus
	NotFound     APIStatus
}

var APIStatusEnum = &apiStatus{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
}

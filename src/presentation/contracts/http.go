package contracts

import "net/http"

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func newHttpResponse(statusCode int, data interface{}) *HttpResponse {
	return &HttpResponse{StatusCode: statusCode, Data: data}
}

func Ok(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusOK, data)
}

func ServerError(err error) *HttpResponse {
	return newHttpResponse(http.StatusInternalServerError, err.Error())
}

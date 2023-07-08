package transport

import "net/http"

type HTTPResponse struct {
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Error   string      `json:"error"`
}

func returnOK(content interface{}) (status int, response HTTPResponse) {
	return http.StatusOK, HTTPResponse{
		Success: true,
		Content: content,
	}
}

func returnErrorResponseFunc(status int, err error) func() (status int, response HTTPResponse) {
	return func() (status int, response HTTPResponse) {
		return status, HTTPResponse{
			Success: false,
			Content: nil,
			Error:   err.Error(),
		}
	}
}

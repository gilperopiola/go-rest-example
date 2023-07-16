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

func returnErrorResponse(status int, err error) (int, HTTPResponse) {
	return status, HTTPResponse{
		Success: false,
		Content: nil,
		Error:   err.Error(),
	}
}

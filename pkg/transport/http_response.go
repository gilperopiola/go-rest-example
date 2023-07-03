package transport

type HTTPResponse struct {
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Error   string      `json:"error"`
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

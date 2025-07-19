package model

type SuccessResponse struct {
	Status     string `json:"status"`
	Detail     string `json:"detail"`
	Data       any    `json:"data"`
	StatusCode int
}

type ErrorResponse struct {
	Status     string `json:"status"`
	Detail     string `json:"detail"`
	Errors     string `json:"error"`
	StatusCode int
}

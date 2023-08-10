package rest

type ErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

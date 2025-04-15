package common

// ErrorResponse — общая ошибка для API
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

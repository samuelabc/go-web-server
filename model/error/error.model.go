package errorModel

type AppError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErrorData string `json:"errorData"`
}

// type ErrorData struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

type ErrorResponse struct {
	Success bool     `json:"success"`
	Error   AppError `json:"error"`
}

package err

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
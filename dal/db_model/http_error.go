package db_model

type HttpErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

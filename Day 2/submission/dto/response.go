package dto

type Response struct {
	StatusCode int         `json:"statusCode" form:"statusCode"`
	Message    string      `json:"message" form:"message"`
	Data       interface{} `json:"data"`
}

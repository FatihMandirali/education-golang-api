package response

type BaseResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message""`
}

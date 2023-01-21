package response

type BaseResponse struct {
	Data       interface{} `json:"data"`
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message""`
}

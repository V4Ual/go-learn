package responses

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func SuccessResponse(statusCode int, msg string, data interface{}) {
	var response APIResponse

	response.StatusCode = statusCode
	response.Message = msg
	response.Data = data
	return

}

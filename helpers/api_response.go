package helpers

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseFailed struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func APIResponseSuccess(message string, data interface{}) ResponseSuccess {
	jsonResponse := ResponseSuccess{
		Message: message,
		Data:    data,
	}
	return jsonResponse
}

func APIResponseSuccessWithoutData(message string) ResponseSuccessWithoutData {
	jsonResponse := ResponseSuccessWithoutData{
		Message: message,
	}
	return jsonResponse
}

func APIResponseFailed(message string) ResponseFailed {
	jsonResponse := ResponseFailed{
		Message: message,
	}
	return jsonResponse
}

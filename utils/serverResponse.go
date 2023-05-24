package utils

type JsonData map[string]interface{}
type ServerResponse struct {
	msg        string
	statusCode int
	status     bool
	data       JsonData
}

func SendResponse(msg string, statusCode int, status bool, data JsonData) ServerResponse {
	response := ServerResponse{
		msg:        msg,
		statusCode: statusCode,
		status:     status,
		data:       data,
	}

	return response
}

func SendSucessResponse(msg string, data ...JsonData) ServerResponse {
	jsonData := JsonData{}

	if len(data) > 0 {
		jsonData = data[0]
	}

	return SendResponse(msg, 200, true, jsonData)
}

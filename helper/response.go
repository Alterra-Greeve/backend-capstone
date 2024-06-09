package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MetadataResponse struct {
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata"`
	Data     interface{} `json:"data"`
}

func FormatResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

func ObjectFormatResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

func MetadataFormatResponse(status bool, message string, metadata interface{}, data interface{}) MetadataResponse {
	response := MetadataResponse{
		Status:   status,
		Message:  message,
		Metadata: metadata,
		Data:     data,
	}
	return response
}

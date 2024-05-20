package helper

func FormatResponse(status bool, message string, data []interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
	return response
}
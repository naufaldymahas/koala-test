package payload

type ResponseFormat struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(s bool, d interface{}) ResponseFormat {
	return ResponseFormat{Status: s, Data: d}
}

func ResponseWithMessage(s bool, m string) ResponseFormat {
	return ResponseFormat{Status: s, Message: m}
}

package response

type MapelResponse struct {
	Status  int      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data`
}

type ResponseDel struct {
	Status  int      `json:"status"`
	Message string      `json:"message"`
}
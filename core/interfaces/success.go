package interfaces

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

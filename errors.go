package cloudElements

// Error is a generic error returned from Cloud Elements, or this client libarary
type Error struct {
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

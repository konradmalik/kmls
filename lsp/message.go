package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// TODO Params
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// TODO Result
	// TODO Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

func NewResponse(id *int) Response {
	return Response{
		RPC: "2.0",
		ID:  id,
	}
}

func NewNotification(method string) Notification {
	return Notification{
		RPC:    "2.0",
		Method: method,
	}
}

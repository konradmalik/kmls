package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... much more here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ServerInfo        `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync *TextDocumentSyncKind `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	textDocumentSyncKind := Full
	return InitializeResponse{
		Response: NewResponse(&id),
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: &textDocumentSyncKind,
			},
			ServerInfo: &ServerInfo{
				Name:    "kmls",
				Version: "0.0.1-dev",
			},
		},
	}
}

type TextDocumentSyncKind int8

const (
	None        TextDocumentSyncKind = 0
	Full        TextDocumentSyncKind = 1
	Incremental TextDocumentSyncKind = 2
)

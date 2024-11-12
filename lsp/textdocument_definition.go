package lsp

type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}

type DefinitionResult struct {
	Contents string `json:"contents"`
}

func NewDefinitionResponse(id int, location Location) DefinitionResponse {
	return DefinitionResponse{
		Response: NewResponse(&id),
		Result:   location,
	}
}

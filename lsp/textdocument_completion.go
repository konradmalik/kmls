package lsp

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string `json:"label"`
	Detail        string `json:"detail"`
	Documentation string `json:"documentation"`
}

func NewCompletionResponse(id int, items []CompletionItem) CompletionResponse {
	return CompletionResponse{
		Response: NewResponse(&id),
		Result:   items,
	}
}

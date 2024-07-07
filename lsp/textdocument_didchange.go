package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VerisonTextDocumentItem          `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	// we don't handle incremental stuff
	Text string `json:"text"`
}

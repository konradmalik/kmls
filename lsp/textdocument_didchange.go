package lsp

type DidChangeNotification struct {
	Notification
	Params DidChangeParams `json:"params"`
}

type DidChangeParams struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	// we don't handle incremental stuff
	Text string `json:"text"`
}

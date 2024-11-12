package analysis

import (
	"fmt"

	"github.com/konradmalik/kmls/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(uri string, position lsp.Position) string {
	return fmt.Sprintf("File: %s, Characters: %d", uri, len(s.Documents[uri]))
}

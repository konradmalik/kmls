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

func (s *State) Definition(uri string, position lsp.Position) lsp.Location {
	return lsp.Location{
		URI: uri,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      position.Line - 1,
				Character: 0,
			},
			End: lsp.Position{
				Line:      position.Line - 1,
				Character: 0,
			},
		},
	}
}

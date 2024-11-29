package analysis

import (
	"fmt"
	"strings"

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

func (s *State) CodeAction(id int, uri string) []lsp.CodeAction {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   lineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*ode with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   lineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*ode",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	return actions
}

func (s *State) Completion(d int, uri string) []lsp.CompletionItem {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "A better VIM",
		},
	}

	return items
}

func lineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}

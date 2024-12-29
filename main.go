package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/konradmalik/kmls/analysis"
	"github.com/konradmalik/kmls/lsp"
	"github.com/konradmalik/kmls/rpc"
)

func main() {
	logger := getLogger("/tmp/kmls.log")

	logger.Println("Starting...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}
}

func writeResponse(writer io.Writer, msg any) error {
	reply := rpc.EncodeMessage(msg)
	_, err := writer.Write([]byte(reply))
	return err
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("initialize: %s", err)
			return
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)

		if err := writeResponse(writer, msg); err != nil {
			logger.Printf("could not respond with InitializeResponse: %s", err)
		}
		logger.Printf("Sent InitializeResponse")

	case "textDocument/didOpen":
		var request lsp.DidOpenNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}

		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		msg := lsp.NewPublishDiagnosticsNotification(lsp.PublishDiagnosticsParams{
			URI:         request.Params.TextDocument.URI,
			Diagnostics: diagnostics,
		})
		if err := writeResponse(writer, msg); err != nil {
			logger.Printf("could not respond with Diagnostics: %s", err)
		}
	case "textDocument/didChange":
		var request lsp.DidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
			msg := lsp.NewPublishDiagnosticsNotification(lsp.PublishDiagnosticsParams{
				URI:         request.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			})
			if err := writeResponse(writer, msg); err != nil {
				logger.Printf("could not respond with Diagnostics: %s", err)
			}
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		content := state.Hover(request.Params.TextDocument.URI, request.Params.Position)
		response := lsp.NewHoverResponse(request.ID, content)
		if err := writeResponse(writer, response); err != nil {
			logger.Printf("could not respond with HoverResponse: %s", err)
		}
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: %s", err)
			return
		}

		location := state.Definition(request.Params.TextDocument.URI, request.Params.Position)
		response := lsp.NewDefinitionResponse(request.ID, location)
		if err := writeResponse(writer, response); err != nil {
			logger.Printf("could not respond with DefinitionResponse: %s", err)
		}
	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/codeAction: %s", err)
			return
		}

		actions := state.CodeAction(request.ID, request.Params.TextDocument.URI)
		response := lsp.NewCodeActionResponse(request.ID, actions)
		if err := writeResponse(writer, response); err != nil {
			logger.Printf("could not respond with CodeActionResponse: %s", err)
		}
	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/completion: %s", err)
			return
		}

		items := state.Completion(request.ID, request.Params.TextDocument.URI)
		response := lsp.NewCompletionResponse(request.ID, items)
		if err := writeResponse(writer, response); err != nil {
			logger.Printf("could not respond with CompletionResponse: %s", err)
		}
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("no good file for a logger")
	}

	return log.New(logfile, "[kmls] ", log.Ldate|log.Ltime|log.Lshortfile)
}

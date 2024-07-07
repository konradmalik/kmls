package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/konradmalik/kmls/lsp"
	"github.com/konradmalik/kmls/rpc"
)

func main() {
	logger := getLogger("/tmp/kmls.log")

	logger.Println("Starting...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("could not parse: %s", err)
			return
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		if _, err := writer.Write([]byte(reply)); err != nil {
			logger.Printf("could not respond with InitializeResponse: %s", err)
		}
		logger.Printf("Sent InitializeResponse")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("could not parse: %s", err)
			return
		}

		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("no good file for a logger")
	}

	return log.New(logfile, "[kmls] ", log.Ldate|log.Ltime|log.Lshortfile)
}

package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var separator = []byte{'\r', '\n', '\r', '\n'}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d%s%s", len(content), separator, content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := splitMessage(msg)
	if !found {
		return "", nil, errors.New("Did not find separator")
	}

	contentLength, err := extractContentLength(header)
	if err != nil {
		return "", nil, err
	}

	var actualContent = content[:contentLength]
	var baseMessage BaseMessage
	if err := json.Unmarshal(actualContent, &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, actualContent, nil
}

// bufio scanner split implementation
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := splitMessage(data)
	if !found {
		// no error, just not ready to split yet
		return 0, nil, nil
	}

	contentLength, err := extractContentLength(header)
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		// also not ready
		return 0, nil, nil
	}

	totalLength := len(header) + len(separator) + contentLength
	return totalLength, data[:totalLength], nil
}

func splitMessage(data []byte) (header []byte, content []byte, found bool) {
	return bytes.Cut(data, separator)
}

func extractContentLength(header []byte) (int, error) {
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, err
	}

	return contentLength, nil
}

package rpc_test

import (
	"testing"

	"github.com/konradmalik/kmls/rpc"
	"github.com/konradmalik/kmls/tests"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	tests.AssertStrings(t, expected, actual)
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)

	tests.AssertNoError(t, err)
	tests.AssertStrings(t, "hi", method)
	tests.AssertInt(t, 15, contentLength)
}

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

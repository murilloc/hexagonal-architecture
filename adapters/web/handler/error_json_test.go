package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello JSON"
	result, _ := jsonError(msg)
	require.Equal(t, `{"message":"Hello JSON"}`, string(result))

}

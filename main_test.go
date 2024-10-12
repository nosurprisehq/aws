package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func TestHandlerFound(t *testing.T) {
	t.Parallel()

	req := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		Path:       "/hello",
	}
	resp, err := handler(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "Hello, World!", resp.Body)
}

func TestHandlerNotFound(t *testing.T) {
	t.Parallel()

	req := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "/not-found",
	}
	resp, err := handler(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, "Not found", resp.Body)
}

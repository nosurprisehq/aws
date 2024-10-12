package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	adapterChi "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
)

func buildAdapter() *adapterChi.ChiLambdaV2 {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()
	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte("world")); err != nil {
			logger.Error("Failed to write the response back to the HTTP connection", "error", err.Error())
		}
	})

	return adapterChi.NewV2(r)
}

func main() {
	adapter := buildAdapter()
	lambda.Start(adapter)
}

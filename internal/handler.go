package internal

import (
	"net/http"
	"sebasromero/github.com/receipt-processor/internal/health"
	"sebasromero/github.com/receipt-processor/internal/receipt"
)

func MainHandler() *http.ServeMux {
	rootHandler := http.NewServeMux()

	rootHandler.HandleFunc("GET /health", health.Health)
	rootHandler.HandleFunc("POST /receipts/process", receipt.Process)
	rootHandler.HandleFunc("GET /receipts/points", receipt.Points)

	rootHandler.Handle("/api/v1/", http.StripPrefix("/api/v1", rootHandler))
	return rootHandler
}

package internal

import (
	"net/http"

	"github.com/sebasromero/receipt-processor/internal/health"
	"github.com/sebasromero/receipt-processor/internal/receipt"
)

func MainHandler() *http.ServeMux {
	rootHandler := http.NewServeMux()

	rootHandler.HandleFunc("GET /health", health.Health)
	rootHandler.HandleFunc("GET /receipts/{id}/points", receipt.Points)
	rootHandler.HandleFunc("GET /receipts", receipt.Receipts)
	rootHandler.HandleFunc("POST /receipts/process", receipt.Process)

	rootHandler.Handle("/api/v1/", http.StripPrefix("/api/v1", rootHandler))
	return rootHandler
}

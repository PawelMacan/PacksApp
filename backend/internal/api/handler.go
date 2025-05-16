package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"packs/internal/domain"
)

type Handler struct {
	calculator domain.PackCalculatorService
	logger     *slog.Logger
}

func NewHandler(calculator domain.PackCalculatorService, logger *slog.Logger) http.Handler {
	h := &Handler{calculator: calculator, logger: logger}

	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", h.corsMiddleware(h.calculateHandler))

	return mux
}

// corsMiddleware adds CORS headers to the response
func (h *Handler) corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// POST /calculate
func (h *Handler) calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.InputRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result, err := h.calculator.CalculatePacks(req.Amount)
	if err != nil {
		h.logger.Error("calculation failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.logger.Error("failed to write response", "error", err)
	}
}

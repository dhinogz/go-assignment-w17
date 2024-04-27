package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct {
	store Store
}

func NewHandlers(store Store) Handlers {
	return Handlers{
		store: store,
	}
}

func (h *Handlers) handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Healthy")
}

func (h *Handlers) handleRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ctx := r.Context()
	params := RecordParams{}
	records := h.store.GetRecords(ctx, params)
	json.NewEncoder(w).Encode(records)
}

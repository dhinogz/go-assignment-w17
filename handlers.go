package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Handlers struct {
	store    Store
	inMemory map[string]string
}

func NewHandlers(store Store) Handlers {
	return Handlers{
		store:    store,
		inMemory: make(map[string]string),
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

type InMemory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h *Handlers) handleInMemory(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No query param for key")
		return
	}

	value, ok := h.inMemory[key]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Key does not exist")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := InMemory{
		Key:   key,
		Value: value,
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) handleCreateInMemory(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var im InMemory
	err := dec.Decode(&im)
	if err != nil {
		msg := "couldn't decode json"
		http.Error(w, msg, http.StatusBadRequest)
	}

	h.inMemory[im.Key] = im.Value

	json.NewEncoder(w).Encode(im)
}

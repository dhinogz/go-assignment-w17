package handlers

import (
	"fmt"
	"net/http"
)

type InMemory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h *Handlers) HandleInMemory(w http.ResponseWriter, r *http.Request) {
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

	data := InMemory{
		Key:   key,
		Value: value,
	}
	err := writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		internalServerError(w, "could not write JSON", err)
	}
}

func (h *Handlers) HandleCreateInMemory(w http.ResponseWriter, r *http.Request) {
	var data InMemory
	err := readJSON(r, data)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "error reading JSON", err)
	}

	h.inMemory[data.Key] = data.Value

	err = writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		internalServerError(w, "could not write JSON", err)
	}
}

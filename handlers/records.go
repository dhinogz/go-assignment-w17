package handlers

import (
	"net/http"

	"github.com/dhinogz/go-assignment-w17/db"
)

type RecordResponse struct {
	Response
	Records []db.Record `json:"records"`
}

func (h *Handlers) HandleRecords(w http.ResponseWriter, r *http.Request) {
	var params db.RecordParams
	err := readJSON(r, params)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "error reading JSON", err)
		return
	}

	// TODO: validate params
	// Validate date formats
	// Validate fields for empty values

	ctx := r.Context()
	records, err := h.store.GetRecords(ctx, params)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "error getting records", err)
		return
	}

	data := RecordResponse{
		Response{
			Code: 0,
			Msg:  "Success",
		},
		records,
	}

	err = writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		internalServerError(w, "could not write JSON", err)
	}
}

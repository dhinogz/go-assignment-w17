package handlers

import (
	"net/http"
	"time"

	"github.com/dhinogz/go-assignment-w17/db"
)

type RecordResponse struct {
	Response
	Records []db.Record `json:"records"`
}

func (h *Handlers) HandleRecords(w http.ResponseWriter, r *http.Request) {
	input := struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int    `json:"minCount"`
		MaxCount  int    `json:"maxCount"`
	}{}
	err := readJSON(r, &input)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "error reading JSON", err)
		return
	}

	if input.MinCount < 0 {
		errorResponse(w, http.StatusBadRequest, "minCount must be larger or equal to 0", nil)
		return
	}
	if input.MaxCount < 0 {
		errorResponse(w, http.StatusBadRequest, "maxCount must be larger or equal to 0", nil)
		return
	}

	params := db.RecordParams{
		MinCount: input.MinCount,
		MaxCount: input.MaxCount,
	}
	params.StartDate, err = time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "wrong date format", err)
		return
	}
	params.EndDate, err = time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "wrong date format", err)
		return
	}

	ctx := r.Context()
	records, err := h.store.GetRecords(ctx, params)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "error getting records", err)
		return
	}

	if len(records) <= 0 {
		errorResponse(w, http.StatusNotFound, "no records found", err)
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

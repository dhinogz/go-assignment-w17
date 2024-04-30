package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/dhinogz/go-assignment-w17/db"
)

type Handlers struct {
	store    db.Store
	inMemory map[string]string
}

func New(store db.Store, inMemory map[string]string) Handlers {
	return Handlers{
		store:    store,
		inMemory: inMemory,
	}
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	data any,
	headers http.Header,
) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return errors.New("could not marshall data to JSON")
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func readJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		return errors.New("could not read JSON")
	}
	return nil
}

func errorResponse(w http.ResponseWriter, status int, msg string, err error) {
	slog.Error(msg, "err", err)
	resp := Response{
		Code: status,
		Msg:  msg,
	}

	err = writeJSON(w, status, resp, nil)
	if err != nil {
		internalServerError(w, "could not write JSON", err)
	}
}

func internalServerError(w http.ResponseWriter, msg string, err error) {
	slog.Error(msg, "err", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}

package errorsx

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// HTTPError writes a warning or error to the slog logger and writes the error message as plain text to the ResponseWriter
func HTTPError(w http.ResponseWriter, err Error, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode < 500 {
		slog.Warn("%s. Stack trace:\n%s", err.Error(), err.Stack())
	} else {
		slog.Error("%s. Stack trace:\n%s", err.Error(), err.Stack())
	}

	w.Write([]byte(err.Error()))
}

type JSONErrorMessageType struct {
	Message string `json:"message"`
}

// HTTPJSONError writes a warning or error to the slog logger and writes the error message as the Message property a JSONErrorMessageType
func HTTPJSONError(w http.ResponseWriter, err Error, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode < 500 {
		slog.Warn(ErrWithStack(err).Error())
	} else {
		slog.Error(ErrWithStack(err).Error())
	}

	json.NewEncoder(w).Encode(JSONErrorMessageType{err.Error()})
}

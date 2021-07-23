package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RecoverFromPanic will handle panic
func RecoverFromPanic(w http.ResponseWriter) {
	if recoveryMessage := recover(); recoveryMessage != nil {
		ERROR(w, errors.New("recover from panic"))
	}
}

// JSON writes the data into the response writer with a JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR reports the error back to the user within a JSON format
func ERROR(w http.ResponseWriter, err error) {
	// default case
	defaultInsprErr := errors.New("server Error")
	JSON(w, http.StatusInternalServerError, defaultInsprErr)

}

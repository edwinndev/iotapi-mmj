package commons

import (
	"encoding/json"
	"net/http"
)

type CustomResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func sendResponse(w http.ResponseWriter, ok bool, message string, status int, data any) {
	response := CustomResponse{
		Ok:      ok,
		Message: message,
		Data:    data,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func sendResponseMinim(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(data)
}

func ApiOK(w http.ResponseWriter, data any) {
	sendResponse(w, true, "Successful", 200, data)
}

func ApiCreated(w http.ResponseWriter, data any) {
	sendResponse(w, true, "Successful", 201, data)
}

func ApiBadRequest(w http.ResponseWriter, message string) {
	sendResponse(w, false, message, 400, nil)
}

func ApiNotFound(w http.ResponseWriter, message string) {
	sendResponse(w, false, message, 404, nil)
}

func ApiUnauthorized(w http.ResponseWriter, message string) {
	sendResponse(w, false, message, 401, nil)
}

func ApiInternalServerError(w http.ResponseWriter) {
	sendResponse(w, false, "Internal Server Error", 500, nil)
}

func ApiMinim(w http.ResponseWriter, data any) {
	sendResponseMinim(w, data)
}

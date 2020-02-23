package controller

import (
	"net/http"
)

// InitRoutes initializes the controller layer's handlers.
func InitRoutes() {
	http.HandleFunc("/drinklogs", logsHandleFunc)
}

func logsHandleFunc(w http.ResponseWriter, r *http.Request) {
	// TODO: consider these carefully for production
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization")

	switch r.Method {
	case http.MethodGet:
		getLogs(w, r)
	case http.MethodPost:
		postLogs(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

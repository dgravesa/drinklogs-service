package controller

import "net/http"

// InitRoutes initializes the controller layer's handlers.
func InitRoutes() {
	http.HandleFunc("/drinklogs", logsHandleFunc)
}

func logsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getLogs(w, r)
	case http.MethodPost:
		postLogs(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

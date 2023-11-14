package util

import (
	"encoding/json"
	"net/http"
	"os"
)

func Response(w http.ResponseWriter, data interface{}, err error) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  err.Error(),
			"status":   false,
			"hostname": hostname,
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":     data,
			"message":  "success",
			"status":   true,
			"hostname": hostname,
		})
	}
}

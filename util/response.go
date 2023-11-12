package util

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":    data,
			"message": "success",
			"status":  true,
		})
	}
}

package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	js, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}
	w.Write(js)
	return nil
}

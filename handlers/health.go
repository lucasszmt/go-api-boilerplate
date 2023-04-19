package handlers

import (
	"go-api-boilerplate/response"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	healthResponse := map[string]string{"status": "ok"}
	if err := response.JSON(w, 200, healthResponse); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json format"})
	}
}

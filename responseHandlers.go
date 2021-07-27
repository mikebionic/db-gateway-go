package main

import (
	"encoding/json"
	"net/http"
)

type BasicApiResponse struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Total   int         `json:"total"`
	Message string      `json:"message"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload map[string]interface{}) {

	var apiResponse BasicApiResponse

	apiResponse.Data = payload["data"]
	if !isNil(apiResponse.Data) {
		apiResponse.Status = 1
	}

	if !isNil(payload["status"]) {
		apiResponse.Status = payload["status"].(int)
	}

	if !isNil(payload["total"]) {
		apiResponse.Total = payload["total"].(int)
	}

	if !isNil(payload["message"].(string)) {
		apiResponse.Message = payload["message"].(string)
	}

	if isNil(apiResponse.Total) {
		apiResponse.Total = apiResponse.Status
	}

	response, _ := json.Marshal(apiResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]interface{}{"message": message})
}

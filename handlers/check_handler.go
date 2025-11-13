package handlers

import (
	"encoding/json"
	"io"
	"linkcheck/models"
	"linkcheck/services"
	"linkcheck/storage"
	"net/http"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный http метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err!=nil {
		http.Error(w, "Ошибка", http.StatusBadRequest)
		return
	}

	var request models.URLCheckRequest
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Ошибка десериализации", http.StatusBadRequest)
		return
	}

	if len(request.Links) == 0 {
		http.Error(w, "Необходимо указать links", http.StatusBadRequest)
		return
	}

	results := make(map[string]string)
	for _, url := range request.Links {
		results[url] = services.CheckURL(url)
	}

	taskID := storage.AddTasks(request.Links, results)

	response := models.URLCheckResponse{
		ID: taskID,
		Results: results,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
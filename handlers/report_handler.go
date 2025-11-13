package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"linkcheck/models"
	"linkcheck/services"
	"linkcheck/storage"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный http метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка", http.StatusBadRequest)
		return
	}

	var request models.ReportRequest
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Ошибка десериализации", http.StatusBadRequest)
		return
	}

	if len(request.LinksNum) == 0 {
		http.Error(w, "Необходимо указать номера links_num", http.StatusBadRequest)
		return
	}

	tasks := storage.GetTasks(request.LinksNum)
	if len(tasks) == 0 {
		http.Error(w, "По данным links_num ничего не найдено", http.StatusNotFound)
		return
	}

	var reportData []models.ReportData
	for _, task := range tasks {
		for link, status := range task.Results {
			reportData = append(reportData, models.ReportData{
				ID:     task.ID,
				Link:    link,
				Status: status,
			})
		}
	}

	pdfFile, err := services.GeneratePDF(reportData)
	if err != nil {
		http.Error(w, "Ошибка при генерации pdf", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, pdfFile)

	os.Remove(pdfFile)
}

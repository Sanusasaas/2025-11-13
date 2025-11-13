package storage

import (
	"encoding/json"
	"linkcheck/models"
	"log"
	"os"
	"sync"
)

var (
	tasks     = make(map[int]models.URLRecord)
	tasksLock sync.Mutex
	currentID = 0
	dataFile  = "data.json"
	shutdown  = false
)

func LoadTasks() {
	tasksLock.Lock()
	defer tasksLock.Unlock()

	_, err := os.Stat(dataFile)
	if err == nil {
		data, err := os.ReadFile(dataFile)
		if err != nil {
			log.Printf("Ошибка чтения файла data.json %v\n", err)
			return
		}
		var storedTasks map[int]models.URLRecord
		if err := json.Unmarshal(data, &storedTasks); err != nil {
			log.Printf("Ошибка: %v\n", err)
			return
		}

		for id, task := range storedTasks {
			tasks[id] = task
			if id > currentID {
				currentID = id
			}
		}
	}
}

func SaveTasks() {
	tasksLock.Lock()
	defer tasksLock.Unlock()

	data, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("Ошибка сериализации: %v\n", err)
		return
	}
	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		log.Printf("Ошибка при записи данных")
		return
	}
}

func AddTasks(links []string, res map[string]string) int {
	tasksLock.Lock()
	defer tasksLock.Unlock()

	currentID++
	taskID := currentID
	tasks[taskID] = models.URLRecord{
		ID:      taskID,
		Links:   links,
		Results: res,
	}

	if !shutdown {
		go SaveTasks()
	}
	return taskID
}

func GetTasks(linksNum []int) []models.URLRecord {
	tasksLock.Lock()
	defer tasksLock.Unlock()

	var res []models.URLRecord
	for _, num := range linksNum {
		if task, exists := tasks[num]; exists {
			res = append(res, task)
		}
	}
	return res
}
func SetShutdown() {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	shutdown = true
}

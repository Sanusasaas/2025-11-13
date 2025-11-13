package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"linkcheck/handlers"
	"linkcheck/storage"
)

func main() {
	storage.LoadTasks()

	http.HandleFunc("/check", handlers.CheckHandler)
	http.HandleFunc("/report", handlers.ReportHandler)

	srv := &http.Server{Addr: ":8080"}
 
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Закрытие сервака...")

		storage.SetShutdown()

		storage.SaveTasks()

		time.Sleep(2 * time.Second)

		if err := srv.Close(); err != nil {
			log.Printf("Ошибка при закрытии сервера: %v", err)
		}
	}()

	log.Println("Сервер успешно запущен на порту :8080")
	err := srv.ListenAndServe()
	if err != nil && err!= http.ErrServerClosed {
		log.Fatalf("Ошибка: %v", err)
	}
}

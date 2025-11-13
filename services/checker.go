package services

import (
	"net/http"
	"time"
)

func CheckURL(url string) string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "unvailable"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "available"
	}
	return "unvailable"
}

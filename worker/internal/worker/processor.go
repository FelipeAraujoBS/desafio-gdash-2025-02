package worker

import (
	"bytes"
	"log"
	"net/http"
	"go-rabbitmq-worker/internal/config"
)

func ProcessAndSend(body []byte, cfg *config.Config) bool {
	log.Printf("Processing message: %s", string(body))
	return sendToNestJSAPI(body, cfg)
}

func sendToNestJSAPI(body []byte, cfg *config.Config) bool {
	req, err := http.NewRequest("POST", cfg.NestJSAPIURL, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: cfg.HTTPTimeout}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Received non-2xx response from NestJS API: %d", resp.StatusCode)
		return false
	}

	log.Printf("Successfully sent message to NestJS API")
	return true
}

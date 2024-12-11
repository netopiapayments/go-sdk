package netopia

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func sendJSON[T any](url, apiKey string, data interface{}) (*T, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

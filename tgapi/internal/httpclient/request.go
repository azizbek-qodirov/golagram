package httpreq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ctxTimeout = 10

func SendHttpGetReques(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(ctxTimeout) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send message, status: %s", resp.Status)
	}

	return resp, nil
}

func SendHttpPostRequest(url string, body []byte) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(ctxTimeout) * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	r, _ := io.ReadAll(resp.Body)
	fmt.Println(string(r))
	return resp, nil
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	token      string
	httpClient *http.Client
	baseURL    string
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
		baseURL:    "https://api.telegram.org/bot",
	}
}

func (c *Client) makeRequest(method, urlStr string, params interface{}) (*http.Response, error) {
	reqUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	var req *http.Request
	if method == "GET" && params != nil {
		// Handle GET request parameters in URL
		query := reqUrl.Query()
		paramsMap, ok := params.(map[string]interface{})
		if ok {
			for key, value := range paramsMap {
				query.Set(key, fmt.Sprint(value))
			}
		}
		reqUrl.RawQuery = query.Encode()
		req, err = http.NewRequest(method, reqUrl.String(), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create GET request: %w", err)
		}
	} else if method == "POST" && params != nil {
		// Handle POST request with JSON body
		reqBody, err := json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		req, err = http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return nil, fmt.Errorf("failed to create POST request: %w", err)
		}
	} else {
		req, err = http.NewRequest(method, reqUrl.String(), nil)

		if err != nil {
			return nil, fmt.Errorf("failed to create %s request: %w", method, err)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

func (c *Client) SetBotCommands(commands interface{}) error {
	url := fmt.Sprintf("%s%s/setMyCommands", c.baseURL, c.token)

	resp, err := c.makeRequest("POST", url, commands)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !response.Ok {
		return fmt.Errorf("failed to set commands: %s", response.Description)
	}

	return nil
}

func (c *Client) SendMessage(chatID int64, text string, replyToMessageID *int64) error {
	params := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}

	if replyToMessageID != nil {
		params["reply_to_message_id"] = *replyToMessageID
	}

	resp, err := c.makeRequest("POST", fmt.Sprintf("%s%s/sendMessage", c.baseURL, c.token), params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}
	return nil
}

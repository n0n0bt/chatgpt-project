package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"models"
)

// Client holds the configuration for the API client.
type Client struct {
	APIurl     string
	APIkey     string
	HTTPclient http.Client
}

// CreateClient initializes a new API client.
func CreateClient(APIkey, APIurl string) (C Client) {
	C.APIkey = APIkey
	C.APIurl = APIurl
	return C
}

// AskGPTansw sends a request to OpenAI and returns the response.
func (c *Client) AskGPTansw(prompt string) (string, error) {
	params := CreatePromptParams(prompt)
	jsonData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.APIurl, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIkey)

	resp, err := c.HTTPclient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response models.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) > 0 && response.Choices[0].Message.Content != "" {
		return response.Choices[0].Message.Content, nil
	}
	return "No response received", nil
}


func CreatePromptParams(prompt string) models.PromptParams {
	return models.PromptParams{
		Model: "gpt-3.5-turbo", // Ensure you are using a chat-compatible model
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens:   150, // Set according to your requirements
		Temperature: 0.7, // Optional: adjust for variability in responses
	}
}


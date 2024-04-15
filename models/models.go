package main

// Message defines a single message for the chat API.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Choice represents a choice returned by the API.
type Choice struct {
	Text         string
	Index        int
	LogProbs     interface{} // Use interface{} if the type varies
	FinishReason string
}

// Response encapsulates the response from the OpenAI API.
type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

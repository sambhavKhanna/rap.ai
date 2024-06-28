package main

type Message struct {
	Role string 	`json:"role"`
	Content string  `json:"content"`
	Image []string	`json:"image,omitempty"`
}


type LlmPayload struct {
	Messages []Message `json:"messages"`
	Model string `json:"model"`
	Stream bool `json:"stream"`
	Options struct {
		Seed string `json:"seed"`
		Temperature float64 `json:"temperature"`
	} `json:"options"`
}

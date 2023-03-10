package bot

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

// AIUser represent a user of openai, it is stateless
type AIUser struct {
	lastActiveTime time.Time
	messages       []openai.ChatCompletionMessage
}

func (user *AIUser) CallOpenai(msg string) (string, error) {
	m := openai.ChatCompletionMessage{
		Role:    "user",
		Content: msg,
	}
	user.messages = append(user.messages, m)
	user.lastActiveTime = time.Now()
	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: cfg.Openai.Temperature,
		TopP:        1.0,
		N:           1,
		Messages:    user.messages,
	}
	log.Printf("ask openai: %+v", msg)
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Printf("%+v", err)
		user.messages = user.messages[:len(user.messages)-1]
		return "", err
	}
	answer := resp.Choices[0].Message
	user.messages = append(user.messages, answer)
	log.Printf("openai answer: %+v", answer.Content)
	return answer.Content, nil
}

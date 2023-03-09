package bot

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

// AIUser represent a user of openai, it is stateless
type AIUser struct {
	LastActiveTime time.Time
	Messages       []openai.ChatCompletionMessage
}

func (user *AIUser) callOpenai(msg string) (openai.ChatCompletionResponse, error) {
	m := openai.ChatCompletionMessage{
		Role:    "user",
		Content: msg,
	}
	user.Messages = append(user.Messages, m)
	user.LastActiveTime = time.Now()
	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: cfg.Openai.Temperature,
		TopP:        1.0,
		N:           1,
		Messages:    user.Messages,
	}
	log.Printf("call openai %+v", msg)
	resp, err := client.CreateChatCompletion(context.Background(), req)
	return resp, err
}

func (user *AIUser) sendAndSaveMsg(msg string) (string, error) {
	resp, err := user.callOpenai(msg)
	if err != nil {
		log.Printf("%+v", err)
		user.Messages = user.Messages[:len(user.Messages)-1]
		return "", err
	}
	answer := resp.Choices[0].Message
	user.Messages = append(user.Messages, answer)
	return answer.Content, nil
}

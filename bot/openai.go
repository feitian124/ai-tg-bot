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

func (user *AIUser) sendAndSaveMsg(msg string) (string, bool, error) {
	ccm := openai.ChatCompletionMessage{
		Role:    "user",
		Content: msg,
	}
	user.Messages = append(user.Messages, ccm)
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
	if err != nil {
		log.Print(err)
		user.Messages = user.Messages[:len(user.Messages)-1]
		return "", false, err
	}

	answer := resp.Choices[0].Message

	user.Messages = append(user.Messages, answer)

	var contextTrimmed bool
	if resp.Usage.TotalTokens > 3500 {
		user.Messages = user.Messages[1:]
		contextTrimmed = true
	}

	return answer.Content, contextTrimmed, nil
}

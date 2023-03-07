package bot

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

// AIUser represent a user of openai
type AIUser struct {
	TelegramID     int64
	LastActiveTime time.Time
	HistoryMessage []openai.ChatCompletionMessage
	LatestMessage  echotron.Message
}

func resetUser(userID int64) {
	delete(users, userID)
}

func (user *AIUser) sendAndSaveMsg(msg string) (string, bool, error) {
	user.HistoryMessage = append(user.HistoryMessage, openai.ChatCompletionMessage{
		Role:    "user",
		Content: msg,
	})
	user.LastActiveTime = time.Now()

	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: cfg.Openai.Temperature,
		TopP:        1.0,
		N:           1,
		Messages:    user.HistoryMessage,
	}

	log.Printf("call openai %+v", req)

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Print(err)
		user.HistoryMessage = user.HistoryMessage[:len(user.HistoryMessage)-1]
		return "", false, err
	}

	answer := resp.Choices[0].Message

	user.HistoryMessage = append(user.HistoryMessage, answer)

	var contextTrimmed bool
	if resp.Usage.TotalTokens > 3500 {
		user.HistoryMessage = user.HistoryMessage[1:]
		contextTrimmed = true
	}

	return answer.Content, contextTrimmed, nil
}

func (user *AIUser) clearUserContextIfExpires() bool {
	if user != nil &&
		user.LastActiveTime.Add(time.Duration(cfg.Openai.IdleTimeout)*time.Second).Before(time.Now()) {
		resetUser(user.TelegramID)
		return true
	}

	return false
}

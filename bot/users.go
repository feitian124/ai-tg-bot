package bot

import (
	"github.com/sashabaranov/go-openai"
	"time"
)

var users = make(map[int64]*User)

func clearUserContextIfExpires(userID int64) bool {
	user := users[userID]
	if user != nil &&
		user.LastActiveTime.Add(time.Duration(cfg.ConversationIdleTimeoutSeconds)*time.Second).Before(time.Now()) {
		resetUser(userID)
		return true
	}

	return false
}

func handleUserPrompt(userID int64, msg string) (string, bool, error) {
	clearUserContextIfExpires(userID)
	var user *User
	var ok bool
	if user, ok = users[userID]; !ok {
		user = &User{
			TelegramID:     userID,
			LastActiveTime: time.Now(),
			HistoryMessage: []openai.ChatCompletionMessage{},
		}
	}
	return user.sendAndSaveMsg(msg)
}

package bot

import (
	"time"
)

var users = make(map[int64]*User)

func (user *User) clearUserContextIfExpires() bool {
	if user != nil &&
		user.LastActiveTime.Add(time.Duration(cfg.Openai.IdleTimeout)*time.Second).Before(time.Now()) {
		resetUser(user.TelegramID)
		return true
	}

	return false
}

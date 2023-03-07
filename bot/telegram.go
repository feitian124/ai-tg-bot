package bot

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

// Recursive type definition of the TGUser state function.
type stateFn func(*echotron.Update) stateFn

// TGUser represent a user of telegram
type TGUser echotron.User

// Bot represent a user of both telegram and openai, and as middle man exchange information between the 2 platform.
type Bot struct {
	TGUser *TGUser
	AIUser *AIUser
	chatID int64
	state  stateFn
	echotron.API
}

var commands = []echotron.BotCommand{
	{Command: "new", Description: "New conversation"},
	{Command: "help", Description: "Show help"},
}

func NewBot(chatID int64) echotron.Bot {
	bot := &Bot{
		chatID: chatID,
		API:    echotron.NewAPI(cfg.TG.Token),
	}
	// We set the default state to the TGUser.handleMessage method.
	bot.state = bot.handleMessage
	_, err := bot.SetMyCommands(nil, commands...)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	return bot
}

func (b *Bot) Update(update *echotron.Update) {
	// Here we execute the current state and set the next one.
	b.state = b.state(update)
}

func (b *Bot) handleMessage(update *echotron.Update) stateFn {
	msg := &Message{update.Message}
	if msg.IsCommand() {
		return b.handleCommand(msg)
	}

	var user *AIUser
	var ok bool
	if user, ok = users[update.Message.From.ID]; !ok {
		user = &AIUser{
			TelegramID:     update.Message.From.ID,
			LastActiveTime: time.Now(),
			HistoryMessage: []openai.ChatCompletionMessage{},
		}
	}

	user.clearUserContextIfExpires()
	answerText, contextTrimmed, err := user.sendAndSaveMsg(update.Message.Text)
	_, err = b.SendMessage(answerText, b.chatID, nil)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	if contextTrimmed {
		_, err := b.SendMessage("Context trimmed.", b.chatID, nil)
		if err != nil {
			log.Printf("error: %+v\n", err)
		}
	}
	return b.handleMessage
}

func (b *Bot) handleCommand(msg *Message) stateFn {
	switch msg.Command() {
	case "new":
		resetUser(msg.From.ID)
		msg.Text = "OK, let's start a new conversation."
	case "help":
		msg.Text = "Write something to start a conversation. Use /new to start a new conversation."
	default:
		msg.Text = "I don't know that command."
	}
	_, err := b.SendMessage(msg.Text, b.chatID, nil)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	return b.handleMessage
}

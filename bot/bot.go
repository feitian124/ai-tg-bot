package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/NicoNex/echotron/v3"
)

// Recursive type definition of the bot state function.
type stateFn func(*echotron.Update) stateFn

type bot struct {
	chatID int64
	state  stateFn
	name   string
	echotron.API
}

var commands = []echotron.BotCommand{
	{Command: "new", Description: "New conversation"},
	{Command: "help", Description: "Show help"},
}

func newBot(chatID int64) echotron.Bot {
	bot := &bot{
		chatID: chatID,
		API:    echotron.NewAPI(cfg.TG.Token),
	}
	// We set the default state to the bot.handleMessage method.
	bot.state = bot.handleMessage
	_, err := bot.SetMyCommands(nil, commands...)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	return bot
}

func (b *bot) Update(update *echotron.Update) {
	// Here we execute the current state and set the next one.
	b.state = b.state(update)
}

func (b *bot) handleMessage(update *echotron.Update) stateFn {
	if strings.HasPrefix(update.Message.Text, "/set_name") {
		_, err := b.SendMessage("Send me my new name!", b.chatID, nil)
		if err != nil {
			log.Printf("error: %+v\n", err)
		}
		// Here we return b.handleName since next time we receive a message it
		// will be the new name.
		return b.handleName
	}

	msg := &Message{update.Message}
	if msg.IsCommand() {
		return b.handleCommand(msg)
	}

	answerText, contextTrimmed, err := handleUserPrompt(update.Message.From.ID, update.Message.Text)
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

func (b *bot) handleName(update *echotron.Update) stateFn {
	b.name = update.Message.Text
	_, err := b.SendMessage(fmt.Sprintf("My new name is %q", b.name), b.chatID, nil)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	// Here we return b.handleMessage since the next time we receive a message
	// it will be handled in the default way.
	return b.handleMessage
}

func (b *bot) handleCommand(msg *Message) stateFn {
	switch msg.Command() {
	case "new":
		resetUser(msg.From.ID)
		msg.Text = "OK, let's start a new conversation."
	case "help":
		msg.Text = "Write something to start a conversation. Use /new to start a new conversation."
	default:
		msg.Text = "I don't know that command"
	}
	_, err := b.SendMessage(msg.Text, b.chatID, nil)
	if err != nil {
		log.Printf("error: %+v\n", err)
	}
	return b.handleMessage
}

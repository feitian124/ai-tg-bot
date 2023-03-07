# openai telegram bot

simple openai ChatGPT Telegram bot with one binary file. 

## Feature

- simple, one binary with no other dependency. 
- well design, makes a new instance of bot for each open chat with a Telegram user, channel or group.
- bot menu supported

## Setup

1. Get your OpenAI API key

   You can create an account on the OpenAI website and [generate your API key](https://platform.openai.com/account/api-keys).

2. Get your telegram bot token

   Create a bot from Telegram [@BotFather](https://t.me/BotFather) and obtain an access token.

3. Install 

```shell
curl -SL https://github.com/feitian124/ai-tg-bot/releases/download/v0.1.0/ai-tg-bot_0.1.0_linux_amd64.tar.gz -o bot.tar.gz && tar -zxvf bot.tar.gz && rm bot.tar.gz
```

## develop

```shell
cp config.yml config.development.yml 
# edit config.development.yml as need, it is ignored by git
# both will load and `config.development.json` will overwrite `config.json`'s configuration

# can restart manually
go run main.go
```
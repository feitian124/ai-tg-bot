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
   Visit github release [page](https://github.com/feitian124/ai-tg-bot/releases/latest), right click the right version to copy its url.
   
   ```shell
   # 1. below url may not up to date, change it to what you copied above
   curl -SLO https://github.com/feitian124/ai-tg-bot/releases/download/v0.1.1/ai-tg-bot_0.1.1_linux_amd64.tar.gz
   
   # 2. unzip files 
   tar -zxvf ai-tg-bot_0.1.1_linux_amd64.tar.gz
   
   # 3. update config.yml as need, then start server
   ./ai-tg-bot
   ```

## develop

   ```shell
   # edit config.development.yml as need, it is ignored by git and it will overwrite `config.json`'s configuration
   cp config.yml config.development.yml 
   
   # if need restart manually
   go run main.go
   ```

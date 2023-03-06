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

3. Install using `go install`

   If you have a Go environment, you can install it with the following command.

4. Install using binary

   You can get prebuilt binaries from [GitHub Releases](https://github.com/feitian124/ai-tg-bot/releases) and put it in `$PATH`

5. Install using Docker-compose(todo)

   Check out [docker-compose.yml](docker-compose.yml) for sample config

6. Set the environment variables and run, see `start.sh`. You should build first `go build -o ai-tg-bot main.go`.

```bash
export OPENAI_KEY=<update_me>
export TELEGRAM_TOKEN=<update_me>
# optional, default is empty. Only allow these users to use the bot. Empty means allow all users.
export ALLOWED_TELEGRAM_ID=
# optional, default is 1.0. Higher temperature means more random responses.
# See https://platform.openai.com/docs/api-reference/chat/create#chat/create-temperature
export MODEL_TEMPERATURE=1.0
# optional, default is 60. Max idle duration for a certain conversation.
# After this duration, a new conversation will be started.
export CONVERSATION_IDLE_TIMEOUT_SECONDS=60

./ai-tg-bot
```

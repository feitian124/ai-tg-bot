#!/usr/bin/env bash

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

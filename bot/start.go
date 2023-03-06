package bot

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/caarlos0/env/v7"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

// maybe changes to https://github.com/jinzhu/configor.git ?

var cfg struct {
	OpenaiKey                      string  `env:"OPENAI_KEY,required"`
	TelegramToken                  string  `env:"TELEGRAM_TOKEN,required"`
	AllowedTelegramID              []int64 `env:"ALLOWED_TELEGRAM_ID" envSeparator:","`
	ConversationIdleTimeoutSeconds int64   `env:"CONVERSATION_IDLE_TIMEOUT_SECONDS,required"`
	ModelTemperature               float32 `env:"MODEL_TEMPERATURE,required"`
}

var client *openai.Client

func Start() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(1)
	}

	client = openai.NewClient(cfg.OpenaiKey)

	log.Printf("%+v\n", cfg)

	// makes a new instance of the struct bot for each open chat with a Telegram user, channel or group.
	dsp := echotron.NewDispatcher(cfg.TelegramToken, newBot)
	log.Printf("bots started ...")
	log.Printf("%+v", dsp.Poll())
}

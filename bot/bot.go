package bot

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/jinzhu/configor"
	"github.com/sashabaranov/go-openai"
	"log"
	"time"
)

var cfg struct {
	TG struct {
		Token      string `required:"true"`
		AllowedIds string
	}
	Openai struct {
		Key         string  `required:"true"`
		Temperature float32 `default:"1.0"`
		IdleTimeout uint    `default:"60"`
	}
}

var client *openai.Client

var dsp *echotron.Dispatcher

func Start() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := configor.New(&configor.Config{Debug: true}).Load(&cfg, "config.yml")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	client = openai.NewClient(cfg.Openai.Key)

	log.Printf("bots started ...")
	// makes a new instance of the struct Bot for each open chat with a Telegram user, channel or group.
	dsp := echotron.NewDispatcher(cfg.TG.Token, NewBot)
	// ensure dispatcher restart poll automatically in case error
	for {
		log.Printf("error: %+v\n", dsp.Poll())
		// In case of connection issues wait 3 seconds before trying to reconnect.
		time.Sleep(3 * time.Second)
	}
}

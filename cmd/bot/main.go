package main

import (
	"flag"

	"github.com/GitH3ll/TTSbot/internal/telegram"
	"github.com/GitH3ll/TTSbot/internal/tts"
)

func main() {
	tokenPtr := flag.String("token", "", "Telegram bot token")
	flag.Parse()
	t1 := tts.NewTextToSpeech("en")
	t2 := tts.NewTextToSpeech("de")

	h, err := telegram.NewHandler(*tokenPtr, map[string]*tts.TextToSpeech{
		"en": t1,
		"de": t2,
	})
	if err != nil {
		panic(err)
	}

	h.Run()
}

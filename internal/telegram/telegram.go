package telegram

import (
	"log"
	"os"

	"github.com/GitH3ll/TTSbot/internal/tts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler struct {
	bot  *tgbotapi.BotAPI
	tts  map[string]*tts.TextToSpeech
	lang string
}

func NewHandler(token string, speech map[string]*tts.TextToSpeech) (*Handler, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Handler{
		bot:  bot,
		tts:  speech,
		lang: "en",
	}, nil
}

func (h *Handler) Run() {
	log.Printf("Authorized on account %s", h.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := h.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "en" || update.Message.Text == "de" {
			h.lang = update.Message.Text
			continue
		}
		t := h.tts[h.lang]
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		err = os.Remove("audio/" + update.Message.Text + ".mp3")
		if err != nil {
			log.Println(err.Error())
		}

		err := t.Convert(update.Message.Text)
		if err != nil {
			log.Println(err.Error())
		}
		msg := tgbotapi.NewAudioUpload(update.Message.Chat.ID, "audio/"+update.Message.Text+".mp3")

		h.bot.Send(msg)
	}
}

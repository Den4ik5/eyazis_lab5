package tts

import htgotts "github.com/hegedustibor/htgo-tts"

type TextToSpeech struct {
	Api *htgotts.Speech
}

func NewTextToSpeech(language string) *TextToSpeech {
	speech := &htgotts.Speech{Folder: "audio", Language: language}
	tts := &TextToSpeech{Api: speech}
	return tts
}

func (t *TextToSpeech) Convert(text string) error {
	err := t.Api.Speak(text)
	if err != nil {
		return err
	}

	return nil
}

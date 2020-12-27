BIN_NAME=ttsbot

.PHONY: build
build:
	go build -o ${BIN_NAME} cmd/bot/main.go


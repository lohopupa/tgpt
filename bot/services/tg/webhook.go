package tg

import (
	logger "bot/common/logger"
	"bot/config"
	"net/http"
	"errors"
	"time"
	"bufio"
	"io"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebHookBot struct {
	instance *tgbotapi.BotAPI
	config config.TgConfig
}

func CreateWH(config config.TgConfig) (*WebHookBot, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, err
	}

	// bot.Debug = true

	logger.Info("Authorized on account %s", bot.Self.UserName)

	go http.ListenAndServeTLS("192.168.0.100:8443", "cert.pem", "key.pem", nil)
	time.Sleep(500 * time.Millisecond)

	certFile := requestFileData{
		FileName: "cert.pem",
	}

	wh, err := tgbotapi.NewWebhookWithCert("https://192.168.0.100:8443/"+bot.Token, certFile)
	if err != nil {
		return nil, err
	}

	_, err = bot.Request(wh)
	if err != nil {
		return nil, err
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return nil, err
	}

	if info.LastErrorDate != 0 {
		logger.Info("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return nil, nil
}

func (bot WebHookBot) Start() error {

	updates := bot.instance.ListenForWebhook("/" + bot.config.ApiKey)

	for update := range updates {
		if update.Message != nil {
			logger.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.instance.Send(msg)
		}
	}
	return nil
}



func (bot WebHookBot) Stop() error {
	return errors.New("Could not stop webhook service")
}

type requestFileData struct {
	FileName string
}

func (t requestFileData) NeedsUpload() bool {
	return true
}

func (t requestFileData) UploadData() (name string, ioOut io.Reader, err error) {
	file, err := os.Open(t.FileName)
	return t.FileName, bufio.NewReader(file), err
}

func (t requestFileData) SendData() string {
	return "ok"
}
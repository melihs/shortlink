// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/bot/slack"
	"github.com/batazor/shortlink/internal/bot/smtp"
	"github.com/batazor/shortlink/internal/bot/telegram"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeFullBotService(ctx context.Context) (*Service, func(), error) {
	bot := InitSlack(ctx)
	telegramBot := InitTelegram(ctx)
	smtpBot := InitSMTP(ctx)
	service, err := NewBotService(bot, telegramBot, smtpBot)
	if err != nil {
		return nil, nil, err
	}
	return service, func() {
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	slack    *slack.Bot
	telegram *telegram.Bot
	smtp     *smtp.Bot
}

// InitSlack - Init slack bot
func InitSlack(ctx context.Context) *slack.Bot {
	slackBot := &slack.Bot{}
	if err := slackBot.Init(); err != nil {
		return nil
	}

	return slackBot
}

// InitTelegram - Init telegram bot
func InitTelegram(ctx context.Context) *telegram.Bot {
	telegramBot := &telegram.Bot{}
	if err := telegramBot.Init(); err != nil {
		return nil
	}

	return telegramBot
}

// InitSMTP - Init SMTP bot
func InitSMTP(ctx context.Context) *smtp.Bot {
	smtpBot := &smtp.Bot{}
	if err := smtpBot.Init(); err != nil {
		return nil
	}

	return smtpBot
}

// FullBotService ======================================================================================================
var FullBotSet = wire.NewSet(InitSlack, InitTelegram, InitSMTP, NewBotService)

func NewBotService(slack2 *slack.Bot, telegram2 *telegram.Bot, smtp2 *smtp.Bot) (*Service, error) {
	return &Service{slack2, telegram2, smtp2}, nil
}

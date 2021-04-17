package service

import (
	"github.com/batazor/shortlink/internal/pkg/notify"
)

type Bot struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

type Service interface {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	Init() error
	Send(message string) error
}

/*
Shortlink application

BFF Web Service
*/
package main

import (
	"os"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
	bff_web_di "github.com/shortlink-org/shortlink/internal/services/bff-web/di"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-bff-web")

	// Init a new service
	service, cleanup, err := bff_web_di.InitializeBFFWebService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Logger.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	handle_signal.WaitExitSignal()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}
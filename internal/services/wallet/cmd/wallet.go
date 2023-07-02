/*
Wallet application

Wallet-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	wallet_di "github.com/shortlink-org/shortlink/internal/services/wallet/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-wallet")

	// Init a new service
	service, cleanup, err := wallet_di.InitializeWalletService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Logger.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}

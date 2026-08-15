package main

import (
	"net/http"
	"os"

	"aifl-backend/factory"
	"aifl-backend/server"
)

func main() {
	appFactory, err := factory.NewFactory()
	if err != nil {
		os.Exit(1)
	}

	routerHandler := appFactory.SetupRouter().(http.Handler)
	appServer := server.NewServer(appFactory.Config, routerHandler)

	if err := appServer.Run(); err != nil {
		os.Exit(1)
	}
}

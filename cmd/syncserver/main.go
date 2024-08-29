package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/kamijoucen/notesync/cmd/syncserver/bootstrap"
	"github.com/kamijoucen/notesync/internal/router"
)

var configFile = flag.String("f", "server.yaml", "config file")

func main() {
	flag.Parse()
	cfg, err := bootstrap.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	router.Init(e, cfg)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}

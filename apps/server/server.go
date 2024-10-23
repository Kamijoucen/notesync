package server

import (
	"flag"
	"fmt"

	"github.com/kamijoucen/notesync/apps/server/internal/router"
	"github.com/labstack/echo/v4"
)

// var configFile = flag.String("f", "server.yaml", "config file")

func main() {
	flag.Parse()
	// cfg, err := bootstrap.LoadConfig(*configFile)
	// if err != nil {
	// 	panic(err)
	// }
	e := echo.New()
	router.Init(e, nil)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))
}

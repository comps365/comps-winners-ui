package main

import (
	"fmt"
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	if err := loadConfig(); err != nil {
		e.Logger.Fatalf("unable to load config")
	}

	pingHandler := handlers.NewPingHandler()

	e.Static("/", "static")
	g := e.Group("/api")
	g.Add(http.MethodGet, "/ping", pingHandler.Ping)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", Port)))
}

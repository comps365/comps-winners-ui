package main

import (
	"fmt"
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/handlers"
	"github.com/comps365/comps-winners-ui/internal/repo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/jmoiron/sqlx"
)

func main() {
	e := echo.New()

	if err := loadConfig(); err != nil {
		e.Logger.Fatalf("unable to load config")
	}

	e.Logger.SetLevel(log.Lvl(LogLvl))

	db, err := sqlx.Connect("mysql", "root:root@(localhost:8889)/guitarcomps")
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo := repo.NewRepo(db, "wp")

	pingHandler := handlers.NewPingHandler()

	e.Static("/", "static")
	g := e.Group("/api")
	g.Add(http.MethodGet, "/ping", pingHandler.Ping)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", Port)))
}

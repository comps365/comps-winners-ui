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
		e.Logger.Fatalf("unable to load config, %v", err)
	}

	e.Logger.SetLevel(log.Lvl(LogLvl))

	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", DBUser, DBPass, DBServer, DBPort, DBName)
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo := repo.NewRepo(db, DBTablePrefix)

	pingHandler := handlers.NewPingHandler()
	lotteryHandler := handlers.NewLotteryHandler(repo)

	e.Static("/", "static")
	g := e.Group("/api")
	g.Add(http.MethodGet, "/ping", pingHandler.Ping)
	g.Add(http.MethodGet, "/lottery/completed", lotteryHandler.GetCompletedLotteries)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", Port)))
}

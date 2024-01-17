package main

import (
	"fmt"
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/handlers"
	"github.com/comps365/comps-winners-ui/internal/repo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/jmoiron/sqlx"
)

func main() {
	e := echo.New()

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
	compHandler := handlers.NewCompetitionHandler(repo)
	authHandler := handlers.NewAuthHandler(SigningKey, AppUser, AppPass)

	e.Static("/", "static")
	g := e.Group("/api")

	authMiddleware := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JWTCustomClaims)
		},
		SigningKey: []byte("secret"),
	})

	g.Add(http.MethodGet, "/ping", pingHandler.Ping, authMiddleware)
	g.Add(http.MethodGet, "/competition/completed", compHandler.GetCompletedCompetitions, authMiddleware)
	g.Add(http.MethodGet, "/competition/:id", compHandler.GetCompetitionDetails, authMiddleware)
	g.Add(http.MethodPost, "/auth/login", authHandler.Login)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", Port)))
}

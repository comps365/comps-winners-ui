package handlers

import (
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/repo"
	"github.com/labstack/echo"
)

type LotteryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LotteryHandler struct {
	repo *repo.Repo
}

func NewLotteryHandler(r *repo.Repo) *LotteryHandler {
	return &LotteryHandler{
		repo: r,
	}
}

func (h *LotteryHandler) GetCompletedLotteries(c echo.Context) error {
	return c.JSON(http.StatusOK, "PONG")
}

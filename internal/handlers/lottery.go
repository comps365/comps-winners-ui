package handlers

import (
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/models"
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
	l, err := h.repo.GetCompletedLotteries()
	if err != nil {
		return notOk(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toLotteriesResponse(l))
}

func toLotteriesResponse(ls []models.Lottery) []LotteryResponse {
	lrs := make([]LotteryResponse, len(ls))
	for i, l := range ls {
		lrs[i] = LotteryResponse{
			Id:   l.Id,
			Name: l.Name,
		}
	}
	return lrs
}

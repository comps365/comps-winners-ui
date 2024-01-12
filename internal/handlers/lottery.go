package handlers

import (
	"net/http"

	"github.com/comps365/comps-winners-ui/internal/models"
	"github.com/comps365/comps-winners-ui/internal/repo"
	"github.com/labstack/echo"
)

type CompetitionHandler struct {
	repo *repo.Repo
}

func NewCompetitionHandler(r *repo.Repo) *CompetitionHandler {
	return &CompetitionHandler{
		repo: r,
	}
}

type CompetitionResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (h *CompetitionHandler) GetCompletedCompetitions(c echo.Context) error {
	l, err := h.repo.GetCompletedCompetitions()
	if err != nil {
		return notOk(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toCompsResponse(l))
}

func toCompsResponse(ls []models.Competition) []CompetitionResponse {
	lrs := make([]CompetitionResponse, len(ls))
	for i, l := range ls {
		lrs[i] = CompetitionResponse{
			Id:   l.Id,
			Name: l.Name,
		}
	}
	return lrs
}

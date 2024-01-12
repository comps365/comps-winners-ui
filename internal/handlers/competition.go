package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

type CompetitionDetailsResponse struct {
	InstantWins   []string           `json:"instantWins"`
	Entries       []CompetitionEntry `json:"entries"`
	UniqueEntries int                `json:"uniqueEntries"`
}

type CompetitionEntry struct {
	Email  string `json:"email"`
	Ticket string `json:"ticket"`
}

func (h *CompetitionHandler) GetCompletedCompetitions(c echo.Context) error {
	l, err := h.repo.GetCompletedCompetitions()
	if err != nil {
		return notOk(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toCompsResponse(l))
}

func (h *CompetitionHandler) GetCompetitionDetails(c echo.Context) error {
	cID, err := validateIntID(c.Param("id"))
	if err != nil {
		return notOk(c, http.StatusBadRequest, err)
	}

	iws, err := h.repo.GetCompetitionInstantWinTickets(cID)
	if err != nil {
		return notOk(c, http.StatusInternalServerError, err)
	}

	entries, err := h.repo.GetEntriesForCompetition(cID)
	if err != nil {
		return notOk(c, http.StatusInternalServerError, err)
	}

	iwsMap := toStrMap(iws)
	ceResps := make([]CompetitionEntry, 0)
	uniquesMap := make(map[string]bool, 0)
	for _, e := range entries {
		for _, t := range e.Tickets {
			if _, ok := iwsMap[t]; ok {
				// ignore tickets that are instant wins, we can't redraw them
				continue
			}
			ceResps = append(ceResps, CompetitionEntry{
				Email:  e.Email,
				Ticket: t,
			})

			if _, ok := uniquesMap[e.Email]; !ok {
				uniquesMap[e.Email] = true
			}
		}
	}

	return c.JSON(http.StatusOK, &CompetitionDetailsResponse{
		InstantWins:   iws,
		Entries:       ceResps,
		UniqueEntries: len(uniquesMap),
	})
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

func validateIntID(id string) (int, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return -1, fmt.Errorf("unable to parse %s, %v", id, err)
	}

	if i < 1 {
		return -1, fmt.Errorf("%d is not a valid ID", i, err)
	}

	return i, nil
}

func toStrMap(s []string) map[string]bool {
	m := make(map[string]bool, len(s))
	for _, ss := range s {
		m[ss] = true
	}
	return m
}

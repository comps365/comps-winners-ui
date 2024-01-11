package repo

import (
	"fmt"

	"github.com/comps365/comps-winners-ui/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db          *sqlx.DB
	tablePrefix string
}

func NewRepo(d *sqlx.DB, t string) *Repo {
	return &Repo{
		db:          d,
		tablePrefix: t,
	}
}

func (r *Repo) GetCompletedLotteries() ([]models.Lottery, error) {
	lotteries := []models.Lottery{}
	if err := r.db.Select(&lotteries, buildCompletedLotteriesQuery(r.tablePrefix)); err != nil {
		return nil, fmt.Errorf("unable to get lotteries, %w", err)
	}
	return lotteries, nil
}

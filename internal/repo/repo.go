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
	var ls []models.Lottery
	if err := r.db.Select(&ls, buildCompletedLotteriesQuery(r.tablePrefix)); err != nil {
		return nil, fmt.Errorf("unable to get lotteries, %w", err)
	}
	return ls, nil
}

func (r *Repo) GetLotteryInstantWinTickets(lID int) ([]int, error) {
	var iws []int
	if err := r.db.Select(&iws, buildInstantWinsQuery(r.tablePrefix), lID); err != nil {
		return nil, fmt.Errorf("unable to get instant win tickets for lottery %d, %v", lID, err)
	}
	return iws, nil
}

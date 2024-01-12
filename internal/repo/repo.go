package repo

import (
	"fmt"
	"regexp"
	"strings"

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

func (r *Repo) GetCompletedCompetitions() ([]models.Competition, error) {
	type compRow struct {
		Id   int
		Name string `db:"lottery_name"`
	}

	var cs []compRow
	if err := r.db.Select(&cs, buildCompletedLotteriesQuery(r.tablePrefix)); err != nil {
		return nil, fmt.Errorf("unable to get lotteries, %w", err)
	}

	comps := make([]models.Competition, len(cs))
	for i, c := range cs {
		comps[i] = models.Competition{
			Id:   c.Id,
			Name: c.Name,
		}
	}

	return comps, nil
}

func (r *Repo) GetCompetitionInstantWinTickets(lID int) ([]int, error) {
	var iws []int
	if err := r.db.Select(&iws, buildInstantWinsQuery(r.tablePrefix), lID); err != nil {
		return nil, fmt.Errorf("unable to get instant win tickets for lottery %d, %w", lID, err)
	}
	return iws, nil
}

func (r *Repo) GetEntriesForCompetition(cID int) ([]models.CompetitionEntry, error) {
	type compEntryRow struct {
		Email   string
		Entries string
	}

	var ers []compEntryRow
	if err := r.db.Select(&ers, buildEntriesQuery(r.tablePrefix), cID); err != nil {
		return nil, fmt.Errorf("unable to get entries for competition %d, %w", cID, err)
	}

	ces := make([]models.CompetitionEntry, len(ers))
	for i, er := range ers {
		ces[i] = models.CompetitionEntry{
			Email:   er.Email,
			Tickets: parseInput(er.Entries),
		}
	}

	return ces, nil
}

func parseInput(input string) []string {
	// example input: a:1:{i:0;i:120;}
	re := regexp.MustCompile(`i:\d+;`)
	matches := re.FindAllString(input, -1)

	var result []string
	for i, match := range matches {
		// we only care about the actual elements, not the indices
		if i%2 != 0 {
			// split on ":" and take the second element
			split := strings.Split(match, ":")
			if len(split) > 1 {
				// remove trailing ; from the element
				element := strings.TrimSuffix(split[1], ";")
				result = append(result, element)
			}
		}
	}

	return result
}

package models

type Competition struct {
	Id   int
	Name string `db:"lottery_name"`
}

type CompetitionEntry struct {
	Email   string
	Tickets []string
}

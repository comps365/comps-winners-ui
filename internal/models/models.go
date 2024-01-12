package models

type Competition struct {
	Id   int
	Name string `db:"lottery_name"`
}

type CompetitionEntry struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   []string
}

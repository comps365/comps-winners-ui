package models

type Lottery struct {
	Id   int
	Name string `db:"lottery_name"`
}

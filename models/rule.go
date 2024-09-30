package models

type Rule struct {
	ID       string
	Name     string
	Exp      string
	GroupID  string
	Priority int
}

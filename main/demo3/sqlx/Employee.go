package main

type Employee struct {
	ID         uint64  `db:"id"          json:"id"`
	Name       string  `db:"name"        json:"name"`
	Department string  `db:"department"  json:"department"`
	Salary     float64 `db:"salary"      json:"salary"`
}

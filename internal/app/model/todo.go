package model

type Todo struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Completed   bool   `db:"completed"`
}

package model

type Subject struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Teacher string `db:"teacher"`
}

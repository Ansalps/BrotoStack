package repo

import "database/sql"
type Repo struct{
	db *sql.DB
}
func NewUserRepo(db *sql.DB) *Repo{
	return &Repo{
		db: db,
	}
}
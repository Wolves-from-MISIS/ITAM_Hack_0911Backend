package repository

import "database/sql"

type HackLabRepository struct {
	db *sql.DB
}

func NewHackLabRepository(db *sql.DB) *HackLabRepository {
	return &HackLabRepository{
		db: db,
	}
}

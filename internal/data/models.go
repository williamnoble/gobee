package data

import (
	"database/sql"
)

type Models struct {
	Bees BeeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Bees: BeeModel{DB: db},
	}
}

type BeeModel struct {
	DB *sql.DB
}

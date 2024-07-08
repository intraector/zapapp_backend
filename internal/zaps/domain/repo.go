package model

import (
	"database/sql"
)

type IZapsRepo struct {
	DB *sql.DB
	IZapsRepoMethods
}

type IZapsRepoMethods interface {
	Create(Car *Car) error
	Update(Car *Car) error
}

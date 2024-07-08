package repo

import (
	"database/sql"
)

type ZapsRepo struct {
	DB *sql.DB
}

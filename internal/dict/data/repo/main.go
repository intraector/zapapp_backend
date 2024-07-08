package dict_repo

import (
	"database/sql"
)

type Repo struct {
	DB *sql.DB
}

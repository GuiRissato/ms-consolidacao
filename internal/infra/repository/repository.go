package repository

import (
	"database/sql"

	"github.com/GuiRissato/ms-consolidacao/internal/infra/db"
)

type Repository struct{
	dbConn *sql.DB
	*db.Queries
}


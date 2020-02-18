package repository

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

func checkResult(result sql.Result) int64 {
	var lastID int64
	if lid, _ := result.RowsAffected(); lid == 1 {
		lastID, _ = result.LastInsertId()
		return lastID
	}
	log.Warn().Msg("query is failed")
	return lastID
}

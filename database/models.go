package database

import "database/sql"

type Advertisement struct {
	ID       uint
	Title    string
	StartAt  sql.NullTime
	EndAt    sql.NullTime
	Age      sql.NullInt16
	Gender   sql.NullString // M, F
	Country  sql.NullString // ISO 3166-1 country code
	Platform sql.NullString // android, ios, web
}

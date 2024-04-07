package models

import "database/sql"

type Advertisement struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	Title   string
	StartAt sql.NullTime
	EndAt   sql.NullTime
}

type Condition struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	Age      sql.NullInt16
	Gender   sql.NullString // M, F
	Country  sql.NullString // ISO 3166-1 country code
	Platform sql.NullString // android, ios, web
}

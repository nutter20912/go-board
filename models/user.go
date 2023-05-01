package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Email           string
	EmailVerifiedAt sql.NullTime
	Password        string
	RememberToken   sql.NullString
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

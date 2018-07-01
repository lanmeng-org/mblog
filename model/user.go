package model

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

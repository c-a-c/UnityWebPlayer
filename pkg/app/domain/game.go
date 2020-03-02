package domain

import (
	"time"
)

type Game struct {
	ID        int
	Title     string
	Auther    string
	CreatedAt *time.Time
	DeletedAt *time.Time
}

type Games []Game

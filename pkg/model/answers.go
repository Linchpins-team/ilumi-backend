package model

import (
	"time"
)

type Answer struct {
	ID      uint
	Author  User
	Mission Mission

	Content string

	CreatedAt time.Time
	UpdatedAt time.Time

	Displayed bool
	Solved    bool
}

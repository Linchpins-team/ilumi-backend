package model

import (
	"time"
)

type User struct {
	ID         uint
	Name       string
	Protrait   string
	Email      string
	JoinedAt   time.Time
	LastOnline time.Time

	// Password is stored in hash
	Password []byte

	Description string
}

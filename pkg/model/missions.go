package model

import (
	"time"
)

type Mission struct {
	ID     uint
	Author User

	Title     string
	Content   string
	Reference string
	Tags      []Tag

	CreatedAt time.Time
	UpdatedAt time.Time

	Displayed bool
}

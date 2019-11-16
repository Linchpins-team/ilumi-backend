package model

import (
	"time"
)

type Mission struct {
	ID     uint
	Author User `gorm:"foreignkey:AuthorID"`
	AuthorID uint

	Title     string
	Content   string
	Reference string
	Tags      []Tag `gorm:"many2many:mission_tags"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Displayed bool
}

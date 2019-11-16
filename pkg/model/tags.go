package model

type Tag struct {
	ID   uint
	Name string

	Missions []Mission `gorm:"many2many:mission_tags"`
}

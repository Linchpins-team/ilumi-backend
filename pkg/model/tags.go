package model

type Tag struct {
	ID   uint
	Name string

	Missions []Mission
}

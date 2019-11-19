package model

import "time"

type AnswerReply struct {
	ID uint

	Replier   User `gorm:"foreignkey:ReplierID"`
	ReplierID uint

	Answer Answer

	CreatedAt time.Time

	Content string
}

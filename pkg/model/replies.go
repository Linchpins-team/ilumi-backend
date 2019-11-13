package model

import "time"

type AnswerReplies struct {
	ID     uint
	User   User
	Answer Answer

	CreatedAt time.Time

	Content string
}

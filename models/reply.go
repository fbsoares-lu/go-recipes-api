package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	Answer   string
	ThreadID int
	Thread   Thread
	UserID   int
	User     User
}

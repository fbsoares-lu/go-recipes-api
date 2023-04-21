package models

type Reply struct {
	Answer   string
	ThreadID int
	Thread   Thread
	UserID   int
	User     User
}

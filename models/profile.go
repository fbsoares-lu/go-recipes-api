package models

type Profile struct {
	Rg        string
	Cpf       string
	Birthday  string
	Biography string
	UserID    int
	User      User
	File      File
}

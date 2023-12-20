package models

type Book struct {
	Author		string		`json:"author"`
	Title		string		`json:"title"`
	Publisher	string		`json:"publisher"`
}
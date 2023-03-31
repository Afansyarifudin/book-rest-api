package model

type Book struct {
	ID       int    `json:"id"`
	Bookname string `json:"book_name"`
	Author   string `json:"author"`
}

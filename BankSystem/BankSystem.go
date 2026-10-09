package main

import (
	"fmt"
)

type Book struct {
	Title    string
	Author   string
	Pages    int
	Borrowed bool
	Year     int
}
type Library struct {
	Books []Book
}
type ErrorBookNotFound string

func (book Book) Strings() string {

	return fmt.Sprintf("\"%d\" by %s (%d pages)", book.Year, book.Author, book.Pages)
}
func (L *Library) AddBook(b Book) {
	NewBook := append(L.Books, b)
	L.Books = NewBook
}
func (f *Library) FindByTitle(title string) (Book, bool) {
	for _, v := range f.Books {
		if title == v.Title {
			fmt.Println("The book Found", v.Strings())
		} else {
			fmt.Println("the book isn't found", false)
		}

	}
}
func main() {
}

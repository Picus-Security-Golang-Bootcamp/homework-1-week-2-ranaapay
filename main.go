package main

import (
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Name        string
	Author      string
	PublishTime string
}

type Library struct {
	BookList *[]Book
}

//insertBooks create static book list and assign it to BookList in Library.
func (l *Library) insertBooks() {
	var bookList []Book
	var bookOne = Book{
		Name:        "In Search of Lost Time",
		Author:      "Marcel Proust",
		PublishTime: "1913",
	}
	var bookTwo = Book{
		Name:        "Ulysses",
		Author:      "James Joyce",
		PublishTime: "1920",
	}
	var bookThree = Book{
		Name:        "Crime and Punishment",
		Author:      "Fyodor Dostoyevski",
		PublishTime: "1866",
	}
	bookList = append(bookList, bookOne, bookTwo, bookThree)
	l.BookList = &bookList
}

//WriteBookList returns console to book list in Library
func (l *Library) writeBookList() {
	for i, book := range *l.BookList {
		fmt.Printf("Book : %d ---------------\nName : %s\nAuthor : %s\nPublish Year : %s\n\n", i+1, book.Name, book.Author, book.PublishTime)
	}
}

//FindBook check is book in the book list and returns the book
func (l *Library) findBook(searchBook string) (Book, bool) {
	searchBook = strings.ToLower(searchBook)
	for _, book := range *l.BookList {
		bLow := strings.ToLower(book.Name)
		if searchBook == bLow {
			return book, false
		}
	}
	return Book{}, false
}

func main() {
	var library Library
	library.insertBooks()
	args := os.Args[1:]
	command := args[0]
	switch command {
	case "list":
		library.writeBookList()
	case "search":
		searchingText := args[1:]
		searchingBook := strings.Join(searchingText, " ")
		book, res := library.findBook(searchingBook)
		if res == false {
			fmt.Printf("Book with %s name is not contain in the book list.", searchingBook)
		} else {
			fmt.Printf("Name : %s\nAuthor : %s\nPublish Year : %s\n\n", book.Name, book.Author, book.PublishTime)
		}
	default:
		fmt.Println("Please write valid command. ( list or search )")
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

var books = []string{
	"In Search of Lost Time",
	"Ulysses",
	"Don Quixote",
	"One Hundred Years of Solitude",
	"The Great Gatsby",
	"War and Peace",
	"Hamlet",
	"The Odyssey",
	"Madame Bovary",
	"The Divine Comedy",
	"Lolita",
	"The Brothers Karamazov",
	"Crime and Punishment",
}

func main(){
	args := os.Args[1:]
	command := args[0]
	switch command {
	case "list":
		WriteBookList()
	case "search":
		searchingText := args[1:]
		searchingBook := strings.Join(searchingText, " ")
		result := FindBook(searchingBook)
		if result == "" {
			fmt.Printf("Book with %s name is not contain in the book list.", searchingBook)
		} else {
			fmt.Println(result)
		}
	default:
		fmt.Println("Please write valid command. ( list or search )")
	}
}

//WriteBookList returns console to book list
func WriteBookList(){
	for _, book := range books {
		fmt.Println(book)
	}
}

//FindBook check is book in the book list and returns the name
func FindBook(searchBook string) string {
	searchBook = strings.ToLower(searchBook)
	for _, b := range books {
		bLow := strings.ToLower(b)
		if searchBook == bLow {
			return b
		}
	}
	return ""
}
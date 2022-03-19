package booksAction

import (
	"fmt"
	"strings"
)

const search = "search"
const list = "list"

type Book struct {
	Id          int
	PageNumber  int
	StockNumber int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	Author
}

type Author struct {
	Name    string
	Surname string
}

// Constructor method for Book struct
func InitBook(id, pageNumber, stockNumber int, price float64, bookName, stockCode, isbn, authorName, authorSurname string) Book {
	book := Book{
		Id:          id,
		PageNumber:  pageNumber,
		StockNumber: stockNumber,
		Price:       price,
		Name:        bookName,
		StockCode:   stockCode,
		Isbn:        isbn,
	}

	book.Author.Name = authorName
	book.Author.Surname = authorSurname

	return book
}

//struct method ????????????????

//Method for perform action according to user input on book list.
func PerformAction(inputList, itemList []string) {
	if len(inputList) == 1 {
		PrintMessagesToConsole()
		return
	}

	firstInput := inputList[1]

	if firstInput == list {
		ListBooks(itemList)
		return
	}

	if firstInput == search {
		if len(inputList) == 2 {
			fmt.Printf("\nPlease enter the book name you want to search!\n\n")
			return
		}
		SearchBooks(itemList, inputList[2:])
		return
	}

	PrintMessagesToConsole()
}

//Method for print the books in the book list.
func ListBooks(books []string) {
	PrintList(books)
}

//Method for search books according to given input and prints the matched books
func SearchBooks(books, keywordList []string) string {

	searchedBook := strings.Join(keywordList, " ")

	for _, book := range books {
		if strings.Contains(strings.ToLower(book), strings.ToLower(searchedBook)) {
			return book
		}
	}

	return "Book not found."
}

//Method for print a list.
func PrintList(list []string) {
	fmt.Println()
	for _, value := range list {
		fmt.Println(value)
	}
	fmt.Println()
}

//Method for print messages to user in the console.
func PrintMessagesToConsole() {
	fmt.Printf("\n--Invalid Input--\n\n")
	fmt.Println("You can use the methods below to make some actions on book list")
	fmt.Println("list: Lists the books")
	fmt.Printf("search \"bookname\": searches the bookname given in the book list\n\n")
}

//Method for check if array contains the item or not.
func Contains(strList []interface{}, str string) bool {
	for _, value := range strList {
		if value == str {
			return true
		}
	}
	return false
}

package main

import (
	"HOMEWORK-1-EZGIUSTUNEL/booksAction"
	"fmt"
	"os"
)

var books []string
var booksInfo []booksAction.Book

func init() {
	//book list
	books = []string{"Simyaci",
		"Bab-i Esrar",
		"Nar Ağaci",
		"Fareler ve İnsanlar",
		"Kürk Mantolu Madonna",
		"Hayvan Çiftliği",
		"Şeker Portakali",
		"Uçurtma Avcisi",
		"Suç ve Ceza",
		"Serenad",
		"Yeraltindan Notlar",
		"Toprak Ana",
		"Fatih Harbiye",
		"Saatleri Ayarlama Enstitüsü",
		"Acimak",
		"Ateşten Gömlek",
		"Çocukluğum",
		"Aşk",
		"Kuyucakli Yusuf",
		"Arkadaş",
	}
}

func main() {
	// var booksInfo []booksAction.Book
	//booksInfo := []booksAction.Book{}
	inputList := os.Args
	booksAction.PerformAction(inputList, books)
	newBook := booksAction.InitBook(5, 10, 30, 44.95, "Nar Ağaci", "asdssds", "11122233", "Nazan", "Demir")
	booksInfo = append(booksInfo, newBook)
	fmt.Println(booksInfo)
}

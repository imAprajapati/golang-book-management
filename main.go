package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

var books = []Book{
	{1, "The Great Gatsby", "F. Scott Fitzgerald"},
	{2, "To Kill a Mockingbird", "Harper Lee"},
	{3, "1984", "George Orwell"},
}

func GetBooks() []Book {
	return books
}

func AddBook() {
	var book Book
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter the author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	book.Id = len(books) + 1
	book.Title = title
	book.Author = author
	books = append(books, book)
	fmt.Println("Book added successfully.")
}

func UpdateBook(id int) {
	for i, book := range books {
		if book.Id == id {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter the title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter the author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			books[i].Title = title
			books[i].Author = author

			fmt.Println("Book updated successfully.")
			return
		}
	}
	fmt.Println("Book not found.")
}

func DeleteBook(id int) {
	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Println("Book deleted successfully.")
			return
		}
	}
	fmt.Println("Book not found.")
}

func menu() {
	fmt.Println("1. List Books")
	fmt.Println("2. Add Book")
	fmt.Println("3. Update Book")
	fmt.Println("4. Delete Book")
	fmt.Println("5. Exit")
}

func main() {
	fmt.Println("Welcome to the Book Store")
	for {
		fmt.Println("############################ Menu ############################")
		menu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			books := GetBooks()
			for _, book := range books {
				fmt.Printf("%d - %s by %s\n", book.Id, book.Title, book.Author)
			}
		case 2:
			AddBook()
		case 3:
			var id int
			fmt.Print("Enter the book id: ")
			fmt.Scanln(&id)
			UpdateBook(id)
		case 4:
			var id int
			fmt.Print("Enter the book id: ")
			fmt.Scanln(&id)
			DeleteBook(id)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

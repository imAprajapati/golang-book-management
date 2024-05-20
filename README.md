# Book Store Management Application

This is a simple Book Store Management application written in Go. It allows users to list, add, update, and delete books from a predefined list. The application is menu-driven and runs in a console environment.

## Features

- **List Books**: Displays a list of all books with their ID, title, and author.
- **Add Book**: Allows the user to add a new book by entering the title and author.
- **Update Book**: Allows the user to update the title and author of an existing book by specifying the book ID.
- **Delete Book**: Allows the user to delete a book by specifying the book ID.

## Usage

1. **List Books**: Press `1` to list all books.
2. **Add Book**: Press `2` to add a new book. You will be prompted to enter the title and author.
3. **Update Book**: Press `3` to update an existing book. You will be prompted to enter the book ID, title, and author.
4. **Delete Book**: Press `4` to delete a book. You will be prompted to enter the book ID.
5. **Exit**: Press `5` to exit the application.

## How to Run

1. **Clone the repository**:
    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2. **Build and Run the application**:
    ```bash
    go run main.go
    ```

3. Follow the on-screen instructions to interact with the application.

## Code Structure

- **Book struct**: Defines the structure of a book with an ID, title, and author.
- **Global variable**: `books` is a slice that stores the list of books.
- **Functions**:
  - `GetBooks()`: Returns the list of books.
  - `AddBook()`: Prompts the user to enter a title and author, then adds a new book to the list.
  - `UpdateBook(id int)`: Prompts the user to update the title and author of a book by its ID.
  - `DeleteBook(id int)`: Deletes a book by its ID.
  - `menu()`: Displays the menu options.
- **main() function**: Runs the application, displaying the menu and handling user input.

## Example

```text
Welcome to the Book Store
############################ Menu ############################
1. List Books
2. Add Book
3. Update Book
4. Delete Book
5. Exit
Enter your choice: 1
1 - The Great Gatsby by F. Scott Fitzgerald
2 - To Kill a Mockingbird by Harper Lee
3 - 1984 by George Orwell
Enter your choice: 2
Enter the title: Brave New World
Enter the author: Aldous Huxley
Book added successfully.
```

## License

This project is licensed under the MIT License.

---

Feel free to contribute to the project by submitting issues or pull requests.

---

If you have any questions or suggestions, please feel free to reach out.

---

Happy coding!

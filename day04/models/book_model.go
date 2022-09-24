package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var Books = []Book{
	{
		ID:     1,
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Year:   1937,
	},
	{
		ID:     2,
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		Year:   1954,
	},
}

func GetAllBooks() []Book {
	return Books
}

func GetBook(id int) *Book {
	for _, book := range Books {
		if book.ID == id {
			return &book
		}
	}

	return &Book{}
}

func CreateBook(book *Book) *Book {
	Books = append(Books, *book)
	return book
}

func UpdateBook(id int, book *Book) *Book {
	for i, b := range Books {
		if b.ID == id {
			book.ID = id
			Books[i] = *book
			return book
		}
	}

	return book
}

func DeleteBook(id int) *Book {
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return &book
		}
	}

	return &Book{}
}

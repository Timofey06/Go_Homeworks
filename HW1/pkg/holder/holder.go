package holder

import "hw1/pkg/book"

type Holder interface {
	AddBook(id int, book book.Book)
	RemoveBook(id int)
	FindBookById(id int) *book.Book
}

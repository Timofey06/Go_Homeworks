package holder

import "hw1/pkg/book"

type MapHolder struct {
	books map[int]book.Book
}

// Methods
func (h *MapHolder) AddBook(id int, book book.Book) {
	h.books[id] = book
}
func (h *MapHolder) RemoveBook(id int) {
	delete(h.books, id)
}
func (h *MapHolder) FindBookById(id int) *book.Book {
	book, exists := h.books[id]
	if exists {
		return &book
	}
	return nil
}

func NewMapHolder() *MapHolder {
	return &MapHolder{
		books: make(map[int]book.Book),
	}
}

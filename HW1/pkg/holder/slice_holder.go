package holder

import "hw1/pkg/book"

type bookIdPair struct {
	id   int
	book book.Book
}

type SliceHolder struct {
	books []bookIdPair
}

func NewSliceHolder() *SliceHolder {
	return &SliceHolder{
		books: make([]bookIdPair, 0),
	}
}

func (h *SliceHolder) AddBook(id int, b book.Book) {
	h.books = append(h.books, bookIdPair{id: id, book: b})
}

func (h *SliceHolder) RemoveBook(id int) {
	for i, pair := range h.books {
		if pair.id == id {
			h.books[i] = h.books[len(h.books)-1]
			h.books = h.books[:len(h.books)-1]
			return
		}
	}
}

func (h *SliceHolder) FindBookById(id int) *book.Book {
	for _, pair := range h.books {
		if pair.id == id {
			return &pair.book
		}
	}
	return nil
}

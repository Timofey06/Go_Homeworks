package library

import (
	"hw1/pkg/book"
	"hw1/pkg/holder"
)

type Library struct {
	holder      holder.Holder
	idHolder    map[string]int
	idGenerator func() int
}

// Getters
func (l *Library) IdGenerator() func() int {
	return l.idGenerator
}

// Setters
func (l *Library) SetIdGenerator(f func() int) {
	l.idGenerator = f
}

// Methods
func (l *Library) AddBook(book book.Book) {
	id := l.idGenerator()
	l.holder.AddBook(id, book)
	l.idHolder[book.Title()] = id
}
func (l *Library) RemoveBook(title string) {
	id, exists := l.idHolder[title]
	if exists {
		l.holder.RemoveBook(id)
		delete(l.idHolder, title)
	}
}
func (l *Library) FindBookByTitle(title string) *book.Book {
	id, exists := l.idHolder[title]
	if exists {
		return l.holder.FindBookById(id)
	}
	return nil
}
func (l *Library) SwitchHolder(h holder.Holder) {
	l.holder = h
	for k := range l.idHolder {
		delete(l.idHolder, k)
	}

}

func NewLibrary(h holder.Holder, idGen func() int) *Library {
	return &Library{
		holder:      h,
		idHolder:    make(map[string]int),
		idGenerator: idGen,
	}
}

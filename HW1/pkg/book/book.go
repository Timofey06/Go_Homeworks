package book

import "fmt"

type Book struct {
	title      string
	author     string
	data       string
	genre      string
	pagesCount int
}

// Getters
func (b *Book) Title() string {
	return b.title
}
func (b *Book) Author() string {
	return b.author
}
func (b *Book) Data() string {
	return b.data
}
func (b *Book) Genre() string {
	return b.genre
}
func (b *Book) PagesCount() int {
	return b.pagesCount
}

// Setters
func (b *Book) SetTitle(title string) {
	b.title = title
}
func (b *Book) SetAuthor(author string) {
	b.author = author
}
func (b *Book) SetData(data string) {
	b.data = data
}
func (b *Book) SetGenre(genre string) {
	b.genre = genre
}
func (b *Book) SetPagesCount(pagesCount int) {
	b.pagesCount = pagesCount
}

// Methods
func (b *Book) PrintBookInformation() {
	fmt.Println(
		"Title: ", b.title,
		"\nAuthor:", b.author,
		"\nData:  ", b.data,
		"\nGenre: ", b.genre,
		"\nPages: ", b.pagesCount)
}

func NewBook(title, author, data, genre string, pagesCount int) Book {
	return Book{
		title:      title,
		author:     author,
		data:       data,
		genre:      genre,
		pagesCount: pagesCount,
	}
}

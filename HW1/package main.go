package main

import "fmt"

type Book struct {
	title      string
	author     string
	data       string
	genre      string
	pagesCount int
}

// Getters
func (b *Book) Title() string   { return b.title }
func (b *Book) Author() string  { return b.author }
func (b *Book) Data() string    { return b.data }
func (b *Book) Genre() string   { return b.genre }
func (b *Book) PagesCount() int { return b.pagesCount }

// Setters
func (b *Book) SetTitle(title string)        { b.title = title }
func (b *Book) SetAuthor(author string)      { b.author = author }
func (b *Book) SetData(data string)          { b.data = data }
func (b *Book) SetGenre(genre string)        { b.genre = genre }
func (b *Book) SetPagesCount(pagesCount int) { b.pagesCount = pagesCount }

// Methods
func (b *Book) PrintBookInformation() {
	fmt.Println(
		"Title: ", b.title,
		"\nAuthor:", b.author,
		"\nData:  ", b.data,
		"\nGenre: ", b.genre,
		"\nPages: ", b.pagesCount)
}

type IHolder interface {
	AddBook(id int, book Book)
	RemoveBook(id int)
	FindBookById(id int) *Book
}

type MapHolder struct {
	books map[int]Book
}

// Methods
func (h *MapHolder) AddBook(id int, book Book) { h.books[id] = book }
func (h *MapHolder) RemoveBook(id int)         { delete(h.books, id) }
func (h *MapHolder) FindBookById(id int) *Book {
	book, exists := h.books[id]
	if exists {
		return &book
	}
	return nil
}

type Library struct {
	holder      IHolder
	idHolder    map[string]int
	idGenerator func() int
}

// Getters
func (l *Library) IdGenerator() func() int { return l.idGenerator }

// Setters
func (l *Library) SetIdGenerator(f func() int) { l.idGenerator = f }

// Methods
func (l *Library) AddBook(book Book) {
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
func (l *Library) FindBookByTitle(title string) *Book {
	id, exists := l.idHolder[title]
	if exists {
		return l.holder.FindBookById(id)
	}
	return nil
}
func (l *Library) SwitchHolder(h IHolder) { l.holder = h }

func main() {
	books := []Book{
		{title: "Harry Potter", author: "J.K. Rowling", data: "1997", genre: "Fantasy", pagesCount: 309},
		{title: "The Hobbit", author: "J.R.R. Tolkien", data: "1937", genre: "Fantasy", pagesCount: 310},
		{title: "1984", author: "George Orwell", data: "1949", genre: "Dystopian", pagesCount: 328},
	}

	holder := &MapHolder{books: make(map[int]Book)}
	library := &Library{
		holder:   holder,
		idHolder: make(map[string]int),
		idGenerator: func() func() int {
			staticId := 1
			return func() int {
				id := staticId
				staticId++
				return id
			}
		}(),
	}

	for _, book := range books {
		library.AddBook(book)
	}

	fmt.Println("Find Harry Potter:")
	found := library.FindBookByTitle("Harry Potter")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Find The Hobbit:")
	found = library.FindBookByTitle("The Hobbit")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	newId := 100
	library.SetIdGenerator(func() int {
		id := newId
		newId++
		return id
	})

	newBook := Book{title: "Brave New World", author: "Aldous Huxley", data: "1932", genre: "Dystopian", pagesCount: 311}
	library.AddBook(newBook)
	fmt.Println("Find Brave New World:")
	found = library.FindBookByTitle("Brave New World")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	newHolder := &MapHolder{books: make(map[int]Book)}
	library.SwitchHolder(newHolder)

	library.AddBook(Book{title: "Dune", author: "Frank Herbert", data: "1965", genre: "Science Fiction", pagesCount: 412})
	library.AddBook(Book{title: "Fahrenheit 451", author: "Ray Bradbury", data: "1953", genre: "Dystopian", pagesCount: 194})

	fmt.Println("Find Dune:")
	found = library.FindBookByTitle("Dune")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Find Fahrenheit 451:")
	found = library.FindBookByTitle("Fahrenheit 451")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}
}

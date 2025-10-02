package main

import (
	"fmt"
	"hw1/pkg/book"
	"hw1/pkg/holder"
	"hw1/pkg/library"
)

func main() {
	books := []book.Book{
		book.NewBook("Harry Potter", "J.K. Rowling", "1997", "Fantasy", 309),
		book.NewBook("The Hobbit", "J.R.R. Tolkien", "1937", "Fantasy", 310),
		book.NewBook("1984", "George Orwell", "1949", "Dystopian", 328),
	}

	myHolder := holder.NewMapHolder()
	myLibrary := library.NewLibrary(
		myHolder,
		func() func() int {
			staticId := 1
			return func() int {
				id := staticId
				staticId++
				return id
			}
		}(),
	)

	for _, book := range books {
		myLibrary.AddBook(book)
	}

	fmt.Println("Find Harry Potter:")
	found := myLibrary.FindBookByTitle("Harry Potter")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Find The Hobbit:")
	found = myLibrary.FindBookByTitle("The Hobbit")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	newId := 100
	myLibrary.SetIdGenerator(func() int {
		id := newId
		newId++
		return id
	})

	newBook := book.NewBook("Brave New World", "Aldous Huxley", "1932", "Dystopian", 311)
	myLibrary.AddBook(newBook)
	fmt.Println("Find Brave New World:")
	found = myLibrary.FindBookByTitle("Brave New World")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	newHolder := holder.NewSliceHolder()
	myLibrary.SwitchHolder(newHolder)

	myLibrary.AddBook(book.NewBook("Dune", "Frank Herbert", "1965", "Science Fiction", 412))
	myLibrary.AddBook(book.NewBook("Fahrenheit 451", "Ray Bradbury", "1953", "Dystopian", 194))

	fmt.Println("Find Dune:")
	found = myLibrary.FindBookByTitle("Dune")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Find Fahrenheit 451:")
	found = myLibrary.FindBookByTitle("Fahrenheit 451")
	if found != nil {
		found.PrintBookInformation()
	} else {
		fmt.Println("Not found")
	}
}

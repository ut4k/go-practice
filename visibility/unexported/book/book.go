package book

import (
	"fmt"
	"unexported/person"
)

type Book struct {
	Author *person.Person
}

func (b *Book) AuthorName() string {
	return fmt.Sprintf("%s %s", b.Author.FirstName(), b.Author.LastName())
}

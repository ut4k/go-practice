package book

import (
	"testing"
	"unexported/person"
)

func TestBookAuthorName(t *testing.T) {
	author := &person.Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	book := &Book{
		Author: author,
	}

	want := "John Doe"
	got := book.AuthorName()

	if got != want {
		t.Error("failed!", got, want)
	}
}

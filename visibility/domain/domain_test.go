package domain

import "testing"

func TestBookAuthorName(t *testing.T) {
	author := &Person{
		firstName: "John",
		lastName:  "Doe",
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

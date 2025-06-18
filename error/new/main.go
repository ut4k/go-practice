package main

import (
	"errors"
	"fmt"
)

type AuthorID int

func (id AuthorID) Valid() bool {
	return id > 0 // 0以下は無効というルール
}

// Author構造体
type Author struct {
	ID   AuthorID
	Name string
}

// Name() メソッド：Authorの名前を返す
func (a *Author) GetName() string {
	return a.Name
}

// Book構造体：AuthorIDを持っている
type Book struct {
	Title    string
	AuthorID AuthorID
}

// GetAuthor: AuthorIDが無効ならエラー、そうでなければ仮のAuthorを返す
func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		return nil, errors.New("GetAuthor: id is invalid")
	}
	return &Author{ID: id, Name: "John Doe"}, nil
}

func GetAuthorName(b *Book) (string, error) {
	a, err := GetAuthor(b.AuthorID)
	if err != nil {
		return "", fmt.Errorf("GetAuthor: %v", err)
	}
	return a.GetName(), nil
}

func main() {
	// 正常
	book := &Book{Title: "Go in Action", AuthorID: 1}
	name, err := GetAuthorName(book)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Author:", name)
	}
	// 異常
	book2 := &Book{Title: "Advanced Go", AuthorID: 0}
	name2, err := GetAuthorName(book2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Author:", name2)
	}
}

package logic

import (
	"errors"
	"encoding/json"
	"sync"
)

type Book struct {
	BookId string
	Name string
	Num int
	Author string
	PublishDate int64
	BorrowCount int
	lock sync.Mutex
}

func NewBook(bookId, name string, num int, author string, publishDate int64) (book *Book) {
	book = &Book {
		BookId: bookId,
		Name:name, 
		Num:num, 
		Author: author,
		PublishDate: publishDate,
	}

	return
}

func (b *Book) Borrow() (err error){
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.Num <= 0 {
		err = errors.New("book is not enough")
		return
	}

	b.Num = b.Num - 1
	b.BorrowCount = b.BorrowCount + 1
	return
}

func (b *Book) Back() (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.Num = b.Num + 1
	return
}

func (b *Book) Marshal() string {
	data, _ := json.Marshal(b)
	return string(data)
}

func (b *Book) UnMarshal(data string) error {
	return json.Unmarshal([]byte(data), b)
}
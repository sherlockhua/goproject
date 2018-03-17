package logic

import (
	"time"
	"errors"
	"encoding/json"
	"sync"
)

type Book struct {
	BookId string `db:"book_id"`
	Name string `db:"name"`
	Num int	`db:"num"`
	Author string  `db:"author"`
	PublishDate time.Time	`db:"publish_time"`
	BorrowCount int	`db:"borrow_count"`
	lock sync.Mutex
}

func NewBook(bookId, name string, num int, author string, publishDate time.Time) (book *Book) {
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
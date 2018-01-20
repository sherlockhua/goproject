package logic

import (
	"sync"
	"fmt"
)

type BookMgr struct {
	BookList []*Book
	//存储bookId 到借书学生列表的信息
	BookStudentMap map[string][]*Student
	//书籍名字到书籍列表的索引
	BookNameMap map[string][]*Book
	//书籍作者到书籍列表的索引
	BookAuthorMap map[string][]*Book
	lock sync.Mutex
}

func NewBookMgr() (bookMgr*BookMgr) {
	bookMgr = &BookMgr {
		BookStudentMap: make(map[string][]*Student, 16),
		BookNameMap:make(map[string][]*Book, 16),
		BookAuthorMap:make(map[string][]*Book, 16),
	}

	return
}

func (b *BookMgr) AddBook(book *Book) (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	//1. 添加到book列表中
	b.BookList = append(b.BookList, book)

	//2. 更新书籍名字到同一个书籍名字对应的book列表
	bookList, ok := b.BookNameMap[book.Name]
	if !ok {
		var tmp []*Book
		tmp = append(tmp, book)
		b.BookNameMap[book.Name] = tmp
	} else {
		bookList = append(bookList, book)
		b.BookNameMap[book.Name] = bookList
	}
	
	//3. 更新书籍作者到同一个作者对应的book列表
	bookList, ok = b.BookAuthorMap[book.Author]
	if !ok {
		var tmp []*Book
		tmp = append(tmp, book)
		b.BookAuthorMap[book.Author] = tmp
	} else {
		bookList = append(bookList, book)
		b.BookAuthorMap[book.Author] = bookList
	}
	return
}

func (b *BookMgr) SearchByBookName(bookName string) (bookList []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookNameMap[bookName]
	return
}

func (b *BookMgr) SearchByAuthor(Author string) (bookList []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookAuthorMap[Author]
	return
}

func (b *BookMgr) SearchByPushlish(min int64, max int64) (bookList []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for _, v := range b.BookList {
		if v.PublishDate >= min && v.PublishDate <= max {
			bookList = append(bookList, v)
		}
	}
	return
}

func (b *BookMgr) Borrow(student *Student, bookId string) (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	
	var book *Book
	for _, v := range b.BookList {
		if v.BookId == bookId {
			book = v
			break
		}
	}

	if book == nil {
		err = fmt.Errorf("book id[%s] it not exist", bookId)
		return
	}

	err = book.Borrow()
	if err != nil {
		return
	}

	student.AddBook(book)
	return
}
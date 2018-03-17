package logic

import (
	"time"
	"database/sql"
	_"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"sync"
	"fmt"
	"encoding/json"
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
	var bookTmp Book
	err = Db.Get(&bookTmp, "select book_id, name, num, author, publish_time from book where name=? and author=?",
		book.Name, book.Author,
	)

	if err != sql.ErrNoRows && err != nil {
		return
	}

	if err == sql.ErrNoRows {
		_, err = Db.Exec("insert into book(name, num, author, publish_time)values(?,?,?,?)",
		book.Name,book.Num, book.Author, book.PublishDate)
	} else {
		_, err = Db.Exec("update book set num = num + ? where name=? and author=?", 
		book.Num, book.Name, book.Author)
	}
	return
}


func (b *BookMgr) SearchByBookName(bookName string) (bookList []*Book) {
	
	sql := fmt.Sprintf("select book_id, name, author, num, publish_time from book where name like '%%%s%%'", bookName)
	Db.Select(&bookList, sql)
	return
}

func (b *BookMgr) SearchByAuthor(Author string) (bookList []*Book) {
	sql := fmt.Sprintf("select book_id, name, author, num, publish_time from book where author like '%%%s%%'", Author)
	fmt.Println(sql)
	
	Db.Select(&bookList, sql)
	return
}

func (b *BookMgr) SearchByPushlish(min time.Time, max time.Time) (bookList []*Book) {
	sql := fmt.Sprintf("select book_id, name, author, num, publish_time from book where publish_time > ? and publish_time < ?", 
		min, max)
	Db.Select(&bookList, sql)
	return
}

func (b *BookMgr) save() {
	data, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("save failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(BookMgrSavePath, data, 0666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
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
	b.save()

	return
}

func (b *BookMgr)GetTop10() (bookList []*Book) {
	sql := fmt.Sprintf("select book_id, name, author, num, publish_time, borrow_count from book order by borrow_count desc limit 10")
	Db.Select(&bookList, sql)
	return
}
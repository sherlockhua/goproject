package main

import (
	"time"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/sherlockhua/goproject/book_mgr/logic"
	"log"
)


var (
	bookMgr *logic.BookMgr
	studentMgr *logic.StudentMgr
)

func init() {
	bookMgr = logic.NewBookMgr()
	studentMgr = logic.NewStudnetMgr()
	err := logic.InitRedis("127.0.0.1:6379", "")
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	err = logic.InitDb("root:@tcp(10.0.0.200:3306)/book_mgr?parseTime=true")
	if err != nil {
		fmt.Printf("init Db failed, err:%v\n", err)
		return
	}
	fmt.Printf("init redis succ\n")
}

func responseError(w http.ResponseWriter, code int) {
	
	m := make(map[string]interface{}, 16)
	m["code"] = code
	m["message"] = getMessage(code)

	data, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(data)
}

func responseSuccess(w http.ResponseWriter, code int, data interface{}) {
	
	m := make(map[string]interface{}, 16)
	m["code"] = code
	m["message"] = getMessage(code)
	m["data"] = data

	dataByte, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(dataByte)
}

func  addBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	bookId := r.FormValue("book_id")
	name := r.FormValue("name")
	numStr := r.FormValue("num")
	author := r.FormValue("author")
	publishDateStr := r.FormValue("publish")

	num, err := strconv.Atoi(numStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	publishDate, err := time.Parse("2006-01-02", publishDateStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	if len(name) == 0 || len(author) == 0 || len(bookId) == 0 {
		responseError(w, ErrInvalidParameter)
		return
	}

	book := logic.NewBook(bookId, name, num, author, publishDate)

	err = bookMgr.AddBook(book)
	//err = bookMgr.AddBookV2(book)
	if err != nil {
		responseError(w, ErrServerBusy)
		return
	}
	responseSuccess(w, ErrSuccess, nil)
}


func  searchBookName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	
	name := r.FormValue("name")


	if len(name) == 0 {
		responseError(w, ErrInvalidParameter)
		return
	}
	bookList := bookMgr.SearchByBookName(name)
	responseSuccess(w, ErrSuccess, bookList)
}


func  searchAuthor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	
	author := r.FormValue("author")
	if len(author) == 0 {
		responseError(w, ErrInvalidParameter)
		return
	}
	bookList := bookMgr.SearchByAuthor(author)
	responseSuccess(w, ErrSuccess, bookList)
}

func  addStudent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	gradeStr := r.FormValue("grade")
	identify := r.FormValue("identify")
	sexStr := r.FormValue("sex")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	sex, err := strconv.Atoi(sexStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	if len(identify) == 0 || len(name) == 0 || (sex != 0 && sex != 1) {
		responseError(w, ErrInvalidParameter)
		return
	}

	stu := logic.NewStudent(id, name, grade, identify, sex)

	studentMgr.AddStudent(stu)
	responseSuccess(w, ErrSuccess, nil)
}

func  bookBorrow(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	
	sidStr := r.FormValue("sid")
	bid := r.FormValue("bid")
	
	sid, err := strconv.Atoi(sidStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	if len(bid) == 0 {
		responseError(w, ErrInvalidParameter)
		return
	}
	student, err  := studentMgr.GetStudentById(sid)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	err = bookMgr.Borrow(student, bid)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}
	studentMgr.Save()
	responseSuccess(w, ErrSuccess, nil)
}

func  studentBookList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	
	sidStr := r.FormValue("sid")
	
	sid, err := strconv.Atoi(sidStr)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	bookList, err  := studentMgr.GetStudentBorrowsBook(sid)
	if err != nil {
		responseError(w, ErrInvalidParameter)
		return
	}

	
	responseSuccess(w, ErrSuccess, bookList)
}


func  getTop10(w http.ResponseWriter, r *http.Request) {
	bookList := bookMgr.GetTop10()
	responseSuccess(w, ErrSuccess,  bookList)
}

func LogHandle(handle func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return  func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixNano()
		handle(w, r)
		end := time.Now().UnixNano()
		cost := (end - start)/1000
		log.Printf("url:%s cost:%d us agent:%s\n", r.RequestURI, cost, r.UserAgent())
	}
}

func main(){
	http.HandleFunc("/book/add", LogHandle(addBook))
	http.HandleFunc("/book/searchName", LogHandle(searchBookName))
	http.HandleFunc("/book/searchAuthor", LogHandle(searchAuthor))

	http.HandleFunc("/student/add", LogHandle(addStudent))
	http.HandleFunc("/student/booklist", LogHandle(studentBookList))
	http.HandleFunc("/book/borrow", LogHandle(bookBorrow))
	http.HandleFunc("/book/top10", LogHandle(getTop10))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server failed, err:", err)
	}
}

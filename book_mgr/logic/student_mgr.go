package logic

import (
	"fmt"
	"sync"
	"encoding/json"
	"io/ioutil"
)

type StudentMgr struct {
	//学生id对应*student map
	StudentMap map[int]*Student
	lock sync.Mutex
}

func NewStudnetMgr() (*StudentMgr) {
	s :=  &StudentMgr{
		StudentMap:make(map[int]*Student, 16),
	}
	s.load()
	return s
}


func (s *StudentMgr) load() {
	data, err := ioutil.ReadFile(StudentMgrSavePath)
	if err != nil {
		fmt.Printf("load failed, err:%v\n", err)
		return
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	fmt.Printf("load data from disk succ")
}

func (s *StudentMgr) AddStudent(stu *Student) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.StudentMap[stu.Id] = stu
	s.Save()
}


func (s *StudentMgr) Save() {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("save failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(StudentMgrSavePath, data, 0666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func (s *StudentMgr)GetStudentById (id int) (stu *Student, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	stu , ok := s.StudentMap[id]
	if !ok {
		err = fmt.Errorf("student id %d is not exists!", id)
		return
	}

	return
}

func (s *StudentMgr)GetStudentBorrowsBook (id int) (bookList []*Book, err error) {
	stu, err := s.GetStudentById(id)
	if err != nil {
		return
	}

	bookList = stu.GetBookList()
	return
}
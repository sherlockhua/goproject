package logic

import (
	"fmt"
	"sync"
)

type StudentMgr struct {
	//学生id对应*student map
	StudentMap map[int]*Student
	lock sync.Mutex
}

func NewStudnetMgr() (*StudentMgr) {
	return &StudentMgr{
		StudentMap:make(map[int]*Student, 16),
	}
}

func (s *StudentMgr) AddStudent(stu *Student) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.StudentMap[stu.Id] = stu
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
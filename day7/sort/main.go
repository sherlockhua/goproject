package main


import (
	//"sort"
	"math/rand"
	"fmt"
)

type Student struct {
	name string
	age int
	score float32
}

type StudentSlice []*Student

func (p StudentSlice) Len() int {
	return len(p)
}

func (p StudentSlice)Less(i, j int) bool {
	return p[i].score > p[j].score
}

func (p StudentSlice)Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var studentArr StudentSlice
	for i := 0; i < 10; i++ {
		var s = &Student {
			name: fmt.Sprintf("李%d", i),
			age: rand.Intn(100),
			score: rand.Float32()*100,
		}
		studentArr = append(studentArr, s)
	}
	
	//sort.Sort(studentArr)
	bubble_sort(studentArr)
	for i := 0; i < len(studentArr); i++{
		fmt.Printf("%#v\n", studentArr[i])
	}
}
package main

import(
	"fmt"
	"sort"
)

// 38, 1, 4, 5, 40
// 1, 38, 4, 5, 10
// 1, 4, 38, 5, 10
// 1, 4, 5, 38, 10
// 1, 4, 5, 10, 38

func bubble_sort(a []int) {

	for i := len(a)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func basic_sort(a []int) {
	sort.Ints(a)
}

//38, 1, 4, 5, 10
//1, 38, 4, 5, 10
//1, 4, 38, 5, 10
func select_sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i+1; j < len(a);j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

//38, 1, 4, 5, 10
//38
//1, 38, 
//1, 4, 38,
//1, 4, 5, 38,
//1, 4, 5, 10, 38
func insert_sort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i;j > 0;j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}

//50, 100, 4, 5, 10, 50
//10, 100, 4, 5, 38, 50
//10, 38, 4, 5, 100, 50
//10, 4, 38, 5, 100, 50
//10, 4, 5, 38, 100, 50
//5, 4, 10,
//4, 5, 10, 38, 
func partion(a []int, left, right int) int {
	var i = left
	var j = right
	for i < j {
		for j > i && a[j] > a[left] {
			j--
		}
		a[j], a[left] = a[left], a[j]
		for i < j && a[i] < a[left] {
			i++
		}
		a[left], a[i] = a[i], a[left]
		fmt.Println(i)
	}
	return i
}

func qsort(a []int, left, right int) {
	if left >= right {
		return
	}

	mid := partion(a, left, right)
	qsort(a, left, mid-1)
	qsort(a, mid+1, right)
}

func main() {
	a := []int{38, 1, 4, 5, 10}
	//bubble_sort(a)
	//basic_sort(a)
	//select_sort(a)
	//insert_sort(a)
	qsort(a, 0, len(a)-1)
	fmt.Println(a)
}
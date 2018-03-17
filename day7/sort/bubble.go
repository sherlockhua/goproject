package main


type SortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}


func bubble_sort(a SortInterface) {
	
		for i := a.Len() -1; i > 0; i-- {
			for j := 0; j < i; j++ {
				//if a[j] > a[j+1] {
				if a.Less(j+1, j) {
					//a[j], a[j+1] = a[j+1], a[j]
					a.Swap(j, j+1)
				}
			}
		}
	}
	
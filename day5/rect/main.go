package main

import(
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	p1 Point
	p2 Point
}



type RectA struct {
	p1 *Point
	p2 *Point
}

func main(){
	var r1 Rect
	var r2 RectA

	r2.p1 = new(Point)
	var r3 = new(Point)
	var r4 = new(Point)
	r2.p2 = new(Point)

	fmt.Println(r3, r4)
	//r1的内存布局
	fmt.Printf("p1.x addr:%p\n", &r1.p1.x)
	fmt.Printf("p1.y addr:%p\n", &r1.p1.y)
	fmt.Printf("p2.x addr:%p\n", &r1.p2.x)
	fmt.Printf("p2.y addr:%p\n", &r1.p2.y)
	fmt.Println()
	fmt.Println()
		//r2的内存布局
		fmt.Printf("p1.x addr:%p\n", &(r2.p1.x))
		fmt.Printf("p1.y addr:%p\n", &(r2.p1.y))
		fmt.Printf("p2.x addr:%p\n", &(r2.p2.x))
		fmt.Printf("p2.y addr:%p\n", &(r2.p2.y))
		fmt.Printf("p1:%p\n", &r2.p1)
		fmt.Printf("P2:%p\n", &r2.p2)
}
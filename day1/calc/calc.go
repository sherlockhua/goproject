/*
dkdkdd
*/
package calc


func Add( a int,  b int)int {
	return a + b
}

func Sub( a int,  b int)int {
	return a - b
}

func Calc(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
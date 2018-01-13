package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){
	// for i:=101;i<=200;i++{
	// 	if isPrime(i){
	// 		fmt.Println(i)
	// 	}
	// }
	// for i:= 100; i<999;i++{
	// 	if isNarcissistic(i){
	// 		fmt.Println(i)
	// 	}
	// }

	// fmt.Println(factorial(10000))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0;i<10;i++{
		// fmt.Println(rand.Int())
		// fmt.Println(r.Intn(100))
		// fmt.Println(r.Float64())
	}
}

func isPrime(n int)bool{
	if n <=1 {
		return false
	}
	for i:=2; i < n; i ++{
		if n % i == 0 {
			return false
		}
	}
	return true
} 

func isNarcissistic(n int) bool{
	if n < 100 || n > 999{
		return false
	}
	i , j , k := float64(n / 100), float64(n/10 %10) , float64(n %10)
	if math.Pow(i,3)+math.Pow(j,3)+math.Pow(k,3) == float64(n){
		return true
	}
	return false
}

func factorial(n int) int{
	if n <=0{
		return -1
	}
	var cnt int = 1
	for i:=2;i<=n;i++{
		cnt = cnt * i
	}
	return cnt
}

// func factorialRecursion(n int) int{
// 	if n <=0 {
// 		return -1
// 	}
// 	return n * factorialRecursion(n-1)
// }

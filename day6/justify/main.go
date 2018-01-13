package main



import (
	"fmt"
)

func justify(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case int:
			fmt.Printf("第 %d 个参数 is int\n", index)
		case int32:
			fmt.Printf("第 %d 个参数 is int32\n", index)
		case float32:
			fmt.Printf("第 %d 个参数 is float32\n", index)
		
		}
	}
}

func main(){
	var a int
	var b float32
	var c int32
	justify(a, b, c)

	
}
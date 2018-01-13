package main

import(
	"time"
	"fmt"
)

func testTime() {
	for {
		now := time.Now()
		fmt.Printf("type of now is:%T\n", now)

		year := now.Year()
		month := now.Month()
		day := now.Day()

		str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d\n", year, month, day, now.Hour(), now.Minute(), now.Second())
		fmt.Println(str)
		time.Sleep(time.Second)

		fmt.Printf("timestamp:%d\n", now.Unix())
	}
}

func testTimeConst() {
	fmt.Printf("Nanosecond :%d\n", time.Nanosecond)
	fmt.Printf("Microsecond:%d\n", time.Microsecond)
	fmt.Printf("Millisecond:%d\n", time.Millisecond)
	fmt.Printf("second     :%d\n", time.Second)
	fmt.Printf("Minute     :%d\n", time.Minute)
	fmt.Printf("Hour       :%d\n", time.Hour)
}

func testTimeFormat() {
	now := time.Now()
	str := now.Format("2006-01-02 03:04:05")
	fmt.Printf("format result:%s\n", str)
}

func testTimeCost() {
	start := time.Now().UnixNano()
	/*
	业务代码
	*/
	time.Sleep(10*time.Millisecond)
	end := time.Now().UnixNano()
	cost := (end - start)/1000
	fmt.Printf("cost:%dus\n", cost)
}

func main() {
	
	//testTime()
	//testTimeConst()
	//testTimeFormat()
	testTimeCost()
}
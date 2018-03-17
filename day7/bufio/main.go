package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

type RandStr struct {

}

func (r *RandStr) Read(p []byte) (n int, err error) {
	fmt.Printf("Len(p)=%d\n", len(p))
	
	source := "abcdefghijk0123456789"
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(source))
		p[i] = source[index];
	}
	p[32] = '\n'
	return len(p), nil
}


func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Println("input:", line)
	rand.Seed(time.Now().UnixNano())
	var randStr = &RandStr{}
	randReader := bufio.NewReader(randStr)
	lineByte, prefix, _ := randReader.ReadLine()
	fmt.Printf("rand:%s prefix:%v\n", string(lineByte), prefix)
}
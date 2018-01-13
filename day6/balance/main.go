package main 

import (
	"fmt"
	"math/rand"
)

func doBalance(balance Balance, addrList []string) (addr string) {
	return balance.DoBalance(addrList)
}

func main() {
	var addrList []string
	for i := 0; i < 5; i++{
		addr := fmt.Sprintf("%d.%d.%d.%d:8080", 
			rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
		addrList = append(addrList, addr)
	}

	var balanceName string
	 
	fmt.Scanf("%s", &balanceName)

	var balance Balance
	if balanceName == "random" {
		balance = &RandBalance{}
	} else if balanceName == "roundrobin" {
		balance = &RoundBalance{}
	} else {
		balance = &RandBalance{}
	}

	for i := 0; i < 10; i++ {
		addr := doBalance(balance, addrList)
		fmt.Println(addr)
	}

}
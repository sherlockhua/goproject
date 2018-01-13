package main

import (
	"math/rand"
)

type RandBalance struct {

}

func (r *RandBalance) DoBalance(addrList []string) string {
	 l := len(addrList)
	 index := rand.Intn(l)
	 return addrList[index]
}
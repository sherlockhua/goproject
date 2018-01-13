package main

type RoundBalance struct {
	curIndex int
}

func (r *RoundBalance) DoBalance(addrList []string) string {
	l := len(addrList)
	r.curIndex = r.curIndex % l
	addr := addrList[r.curIndex]
	r.curIndex++
	return addr
}

package test

import(
	"testing"
	"code.oldboy.com/day1/calc"
)


func TestAdd(t *testing.T) {

	var sum  int
	sum = calc.Add(5, 6)
	if sum  != 10 {
		t.Fatalf("add is not right, sum:%v expected:11", sum)
	}

	t.Logf("add is ok")
}


func TestSub(t *testing.T) {
	
		var sum  int
		sum = calc.Sub(5, 6)
		if sum  != -1 {
			t.Fatalf("add is not right, sum:%v expected:11", sum)
		}
	
		t.Logf("sub is ok")
	}
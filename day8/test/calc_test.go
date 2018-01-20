package test


import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 8)
	if result != 10 {
		t.Fatalf("add is not right")
		return
	}

	t.Logf("add is right")
}

func TestSub(t *testing.T) {
	result := sub(2, 8)
	if result != -6 {
		t.Fatalf("add is not right")
		return
	}

	t.Logf("add is right")
}
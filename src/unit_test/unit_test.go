package service

import (
	"testing"
)

func TestHelloWorld(t *testing.T){
	t.Log("Enter TestHelloWorld.")
}

func TestAdd(t *testing.T)  {
	r := Add(1, 2)
	if r !=3 {
		t.Errorf("1+2 expected be 3")
	}

}

// go test -v  unit_test.go
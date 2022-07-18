package testlib11

import (
	"fmt"

	"github.com/tyaps/testlib2"
)

func PrintName(s string) {
	testlib2.PrintName("aaa")
}

func PrintAny(s string) {
	fmt.Printf("some print by myself: testlib11")
}

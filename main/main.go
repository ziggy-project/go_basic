package main

import (
	"fmt"

	"github.com/ak-py-proj/go_basic/mymath"
)

func main() {
	if sum := mymath.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}

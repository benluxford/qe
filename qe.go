package main

import (
	"fmt"

	"github.com/benluxford/qe/vector"
)

type Vector []complex128

func main() {
	v := vector.New(1, 2, 3, 4, 5)
	fmt.Println(v)
	c := vector.TensorProduct(v, v)
	fmt.Println(c)
}

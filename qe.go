package main

import (
	"fmt"

	"github.com/benluxford/qe/qubit"
)

func main() {
	// v := vector.New(1, 2, 3, 4, 5)
	// fmt.Println(v)
	// c := vector.TensorProduct(v, v)
	// fmt.Println(c)
	// m := matrix.Matrix{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	// fmt.Println(m)
	// // fmt.Println(m.Transpose())
	// fmt.Println(matrix.TensorProduct(m, m, m))
	// fmt.Println(matrix.Eps(22.5, 15.2))
	q := qubit.New(1, 2, 3, 4, 5)
	fmt.Println(q)
}

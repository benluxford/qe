package main

import (
	"fmt"

	"github.com/benluxford/qe/payments"
	"github.com/benluxford/qe/vector"
)

type Vector []complex128

func main() {
	p := payments.NewResolver("sk_test_4iWkCqj4u6AVaG1pgghOpj4s")
	fmt.Println(p)
	v := vector.New(1, 2, 3, 4, 5)
	fmt.Println(v)
	// s := vector.New(1, 2, 3, 4, 5)
	// a := vector.NewZero(33)
	// fmt.Println(a)
	c := vector.TensorProduct(v, v)
	fmt.Println(c)
	// fmt.Println(complex(1, -5))
}

func examples() {
	//--------------------//
	// 			copy		  //
	//--------------------//
	s := Vector{12345, 54321, 55555, 88888, 99999}
	// copy will take the n elements, meanming that the first 3 will be easy as to get by saying 3
	g := make(Vector, len(s))
	copy(g, s)
	fmt.Println(g)
}

package vector

import (
	"math/cmplx"

	"github.com/benluxford/qe/matrix"
)

// Vector : The vector of a Qubit
type Vector []complex128

// New : Creates a new Vector, returns the slice of complex numbers
func New(values ...complex128) (v Vector) {
	// add all the params to the complex number slice
	for _, val := range values {
		v = append(v, val)
	}
	// return the new vector
	return
}

// NewZero : Returns a new vector with the n number of zeros
func NewZero(n int) (v Vector) {
	// make vector of n length, no contents
	v = make(Vector, n)
	// return the new vector
	return
}

// Clone : Clones the current vector and returns the clone
func (v Vector) Clone() (vClone Vector) {
	// make a new vector at the target length
	vClone = make(Vector, len(v))
	// copy parent vector into clone
	copy(vClone, v)
	// return the cloned vector
	return
}

// Conjugate : Returns the conjugate of the current vector
func (v Vector) Conjugate() (vConj Vector) {
	// for every complex number in the vector
	for _, val := range v {
		// calculate and append the conjugate
		vConj = append(vConj, cmplx.Conj(val))
	}
	// return the inverted vector
	return
}

// Add : Returns the sum of two vectors of the same length
// TODO - this needs to be modified to allow many vectors for addition
// and to resize to the largest vector......
func (v Vector) Add(input Vector) (total Vector) {
	// get the vector length
	vectorLength := len(v)
	// create a vector of the same length
	total = make(Vector, vectorLength)
	// loop over all values and add
	for i := 0; i < vectorLength; i++ {
		total[i] = v[i] + input[i]
	}
	// return sum vectors
	return
}

// Multiply : Returns product of the input applied to all components in vector
func (v Vector) Multiply(input complex128) (total Vector) {
	for _, component := range v {
		total = append(total, component*input)
	}
	return
}

// TensorProduct : Returns the tensor product of two vectors of the same length
// e.g. vectorA{1, 2, 3} * vectorB{1, 2, 3} = vectorC{1, 2, 3, 2, 4, 6, 3, 6, 9}
// TODO - this needs to be modified to allow many vectors
func (v Vector) TensorProduct(input Vector) (total Vector) {
	// loop over all base vector values
	for _, vVal := range v {
		// loop over all input vectors values
		for _, inputVal := range input {
			// append the product of the base vector value and the input vector value
			total = append(total, vVal*inputVal)
		}
	}
	return
}

// InnerProduct : Return the product of the vector values * the conjugate of the input vector
// TODO - resize to the largest vector......
func (v Vector) InnerProduct(input Vector) (product complex128) {
	// get the conjugate of the input
	conjugateInput := input.Conjugate()
	// loop over all values in current vector
	for i := 0; i < len(v); i++ {
		// calculate the product
		product = product + v[i]*conjugateInput[i]
	}
	// return the product
	return
}

// IsOrthogonal : Returns bool if the input vector is perpendicular to the current vector
func (v Vector) IsOrthogonal(input Vector) (orthogonal bool) {
	if v.InnerProduct(input) == complex(0, 0) {
		orthogonal = true
	}
	return
}

// Normalise : Returns the normalised vector as a complex number
func (v Vector) Normalise() (normalised complex128) {
	normalised = cmplx.Sqrt(v.InnerProduct(v))
	return
}

// IsUnit : Returns true if the normalised vector has a length of 1
func (v Vector) IsUnit() (unit bool) {
	if v.Normalise() == complex(1, 0) {
		unit = true
	}
	return
}

// TensorProductN : Returns the tensor product of a given vector n times
func TensorProductN(input Vector, bit ...int) (product Vector) {
	// set the product to be the input
	product = input
	// if n == 0 return the product of vector * 0 = input vector
	if len(bit) < 1 {
		return
	}
	// for the number of iterations, compound the product of the vector
	for i := 1; i < bit[0]; i++ {
		product = product.TensorProduct(input)
	}
	return
}

// TensorProduct : Returns the tensor product of a given set of vectors
func TensorProduct(input ...Vector) (product Vector) {
	// set the product to the first vector
	product = input[0]
	// from the second vector on...
	for i := 1; i < len(input); i++ {
		// calculate the product of the remaining vectors
		product = product.TensorProduct(input[i])
	}
	return
}

// Apply : Return vector multiplied by matrix
func (v Vector) Apply(input matrix.Matrix) (appliedVector Vector) {
	// get the number of rows and columns
	mRows, _ := input.Dimension()
	// for all of the rows in the matrix
	for i := 0; i < mRows; i++ {
		// tmp := complex(0, 0)
		var component complex128
		// for all components in vector
		for j := 0; j < len(v); j++ {
			// add the product of the input by the vector
			component += input[i][j] * v[j]
		}
		// add the component to the vector
		appliedVector = append(appliedVector, component)
	}
	// return the applied vector
	return
}

// Equals : Return if all components of current vector and input vector == 0 or are > e
func (v Vector) Equals(input Vector, eps ...float64) (equals bool) {
	// if the vectors are of diff length, return false
	if len(v) != len(input) {
		return
	}
	// Get the boundary number (probably 0.0)
	e := matrix.Eps(eps...)
	// for each component
	for i := 0; i < len(v); i++ {
		// check that current vector component - input component == 0.0 or is greater than e
		if cmplx.Abs(v[i]-input[i]) > e {
			return
		}
	}
	// nothing was triggered, the vectors are the same :)
	equals = true
	return true
}

// Dimension : Returns the length of the vector
func (v Vector) Dimension() int {
	return len(v)
}

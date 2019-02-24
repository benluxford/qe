package matrix

import "math/cmplx"

// Matrix : A matrix of complex numbers
type Matrix [][]complex128

// Dimension : returns the overall dims of the matrix, height and width
// TODO : save the rows and columns as a value in the Matrix, will enable the removal
// of all row and column counts in other functions
func (m Matrix) Dimension() (rows, columns int) {
	// total number of rows in matrix
	rows = len(m)
	// total number of columns in matrix
	columns = len(m[0])
	return
}

// Transpose : Returns matrix with all 1'st column values within the first row and etc
// e.g. Matrix{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}} => Matrix{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
func (m Matrix) Transpose() (swapped Matrix) {
	// count the number of rows and columns in cuttent matrix
	rows, columns := m.Dimension()
	// for all of the rows
	for i := 0; i < rows; i++ {
		// create a temp row
		v := []complex128{}
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// append values to the temp row
			v = append(v, m[j][i])
		}
		// add temp row to the swapped map
		swapped = append(swapped, v)
	}
	return
}

// Equals : Returns bool, if input matrix matches the current matrix = true
func (m Matrix) Equals(input Matrix, eps ...float64) (match bool) {
	// count the number of rows and columns in cuttent matrix
	mRows, mColumns := m.Dimension()
	// count the number of rows and columns in input matrix
	inputRows, inputColumns := input.Dimension()

	// if the columns or rows do not match return false
	if mRows != inputRows || mColumns != inputColumns {
		return
	}

	// If present, return the first eps value (dont know why this has been done, will have to check)
	e := Eps(eps...)
	// for all of the rows
	for i := 0; i < mRows; i++ {
		// for each column in the current row
		for j := 0; j < mColumns; j++ {
			// if the sum of the corresponding column and number is not 0.0 or lower than e return
			if cmplx.Abs(m[i][j]-input[i][j]) > e {
				return
			}
		}
	}
	// no differences were detected
	match = true
	return
}

// Conjugate : Returns the conjugate of the current matrix
func (m Matrix) Conjugate() (mConj Matrix) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// for all of the rows
	for i := 0; i < rows; i++ {
		// create a temp row
		v := []complex128{}
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// append values to the temp row
			v = append(v, cmplx.Conj(m[i][j]))
		}
		// add temp row to the conjugate matrix
		mConj = append(mConj, v)
	}
	return
}

// Dagger : returns the matrix after Transpose and Conjugate are applied
func (m Matrix) Dagger() (transposeConjugate Matrix) {
	transposeConjugate = m.Transpose().Conjugate()
	return
}

// IsHermite : Returns if the number is Hermite polynomial
func (m Matrix) IsHermite(eps ...float64) (hermite bool) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// dagger the matrix
	dagger := m.Dagger()
	// If present, return the first eps value (dont know why this has been done, will have to check)
	e := Eps(eps...)
	// for all of the rows
	for i := 0; i < rows; i++ {
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// if the sum of the matrix and daggered matrix is not 0.0 or lower than e > return
			if cmplx.Abs(m[i][j]-dagger[i][j]) > e {
				return
			}
		}
	}
	// If not triggered, map is Hermite polynormal
	hermite = true
	return
}

// IsUnitary : Returns if the matrix is unitary perfect number
func (m Matrix) IsUnitary(eps ...float64) (unitary bool) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// get the applied dagger of the matrix
	appliedDagger := m.Apply(m.Dagger())
	// If present, return the first eps value (dont know why this has been done, will have to check)
	e := Eps(eps...)
	// for all of the rows
	for i := 0; i < rows; i++ {
		// for each column in the current row
		for j := 0; j < columns; j++ {
			if i == j {
				// if the modulus of component - 1 > e || > 0.0 if no e => false
				if cmplx.Abs(appliedDagger[i][j]-complex(1, 0)) > e {
					return
				}
				continue
			}
			// if the modulus of component > e || > 0.0 if no e => false
			// strikes me as strings, this could be changed, why subtract 0 from the applied dagger number?
			if cmplx.Abs(appliedDagger[i][j]-complex(0, 0)) > e {
				return
			}
		}
	}
	// nothing hit, the matrix is unitary
	unitary = true
	return
}

// Apply : Return matrix multiplied by another of the same size
// e.g. Matrix{{1, 2, 3},{1, 2, 3},{1, 2, 3}} => {1*1+2*1+3*1}, {1*2+2*2+3*2}, {1*3+2*3+3*3}....
func (m Matrix) Apply(input Matrix) (applied Matrix) {
	// get the number of rows and columns
	inputRows, inputColumns := input.Dimension()
	// for all rows
	for i := 0; i < inputRows; i++ {
		// create the new vector
		vector := []complex128{}
		// for each column in the current row
		for j := 0; j < inputColumns; j++ {
			// create the new component
			var component complex128
			// matrix multiplication (matrix is the same size, no need to use the row count from m)
			for k := 0; k < inputRows; k++ {
				// add the product to component
				component += input[i][k] * m[k][j]
			}
			// append the component to the vector
			vector = append(vector, component)
		}
		// append the vector to the applied matrix
		applied = append(applied, vector)
	}
	// return the applied matrix
	return
}

// Multiply : Returns current matrix with each component multiplied by the input
func (m Matrix) Multiply(input complex128) (multiMatrix Matrix) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// for all rows in the matrix
	for i := 0; i < rows; i++ {
		// create the new vector/row
		vector := []complex128{}
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// append the product to the new vector
			vector = append(vector, input*m[i][j])
		}
		// append the vector/row to the matrix
		multiMatrix = append(multiMatrix, vector)
	}
	// return the "product" matrix
	return
}

// Add : Returns the current matrix with the sum of all input matrix's components applied
func (m Matrix) Add(input Matrix) (sumMatrix Matrix) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// for all the rows
	for i := 0; i < rows; i++ {
		// create the new vector
		vector := []complex128{}
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// append the sum of the two corresponding components from current and input matrix
			vector = append(vector, m[i][j]+input[i][j])
		}
		// append the vector to the "sum" matrix
		sumMatrix = append(sumMatrix, vector)
	}
	// return the sum matrix
	return
}

// Subtract : subtracts the components of the input matrix from the components of the current matrix
func (m Matrix) Subtract(input Matrix) (subMatrix Matrix) {
	// get the number of rows and columns
	rows, columns := m.Dimension()
	// for each row
	for i := 0; i < rows; i++ {
		// create the new vector
		vector := []complex128{}
		// for each column in the current row
		for j := 0; j < columns; j++ {
			// append the current matrix component - the input matrix component
			vector = append(vector, m[i][j]-input[i][j])
		}
		// append the vector to the "sub" matrix
		subMatrix = append(subMatrix, vector)
	}
	// return the current matrix, less the value of the components of the input matrix
	return
}

// Trace : Returns the sum of diagonal across the matrix left to right - top to bottom
func (m Matrix) Trace() (sum complex128) {
	// get the number of rows and columns
	rows, _ := m.Dimension()
	// var sum complex128
	for i := 0; i < rows; i++ {
		sum += m[i][i]
	}
	// return the sum of the traced numbers
	return
}

// TensorProduct : Returns the tensor product of the current and input matrix
func (m Matrix) TensorProduct(input Matrix) (product Matrix) {
	// get the number of rows and columns
	mRows, mColumns := m.Dimension()
	inputRows, inputColumns := input.Dimension()
	// create the temp slice of matrix
	tmp := []Matrix{}
	// for row in current matrix
	for i := 0; i < mRows; i++ {
		// for all columns in the current row
		for j := 0; j < mColumns; j++ {
			// multiply the input matrix by the current matrix component
			// append multiplied matrix to tem slice
			tmp = append(tmp, input.Multiply(m[i][j]))
		}
	}

	// for matrix in tmp matrix slice
	for l := 0; l < len(tmp); l = l + mRows {
		// for rows in each matrix
		for j := 0; j < inputRows; j++ {
			// create the new vector
			vector := []complex128{}
			// for each row in current matrix
			for i := l; i < l+mRows; i++ {
				// for each component in row
				for k := 0; k < inputColumns; k++ {
					// append the value from temp > matrix > row > component
					vector = append(vector, tmp[i][j][k])
				}
			}
			// append the vector to the product matrix
			product = append(product, vector)
		}
	}
	// return the TensorProduct of the two matrix
	return
}

// func TensorProductN(m Matrix, bit ...int) Matrix {
// 	if len(bit) < 1 {
// 		return m
// 	}

// 	m0 := m
// 	for i := 1; i < bit[0]; i++ {
// 		m0 = m0.TensorProduct(m)
// 	}

// 	return m0
// }

// func TensorProduct(m ...Matrix) Matrix {
// 	m0 := m[0]
// 	for i := 1; i < len(m); i++ {
// 		m0 = m0.TensorProduct(m[i])
// 	}

// 	return m0
// }

// Eps : A very strange little function that returns the first float or 0.0
func Eps(eps ...float64) (value float64) {
	if len(eps) > 0 {
		value = eps[0]
	}
	return
}

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

// func (m0 Matrix) IsHermite(eps ...float64) bool {
// 	p, q := m0.Dimension()
// 	m := m0.Dagger()
// 	e := Eps(eps...)

// 	for i := 0; i < p; i++ {
// 		for j := 0; j < q; j++ {
// 			if cmplx.Abs(m0[i][j]-m[i][j]) > e {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func (m0 Matrix) IsUnitary(eps ...float64) bool {
// 	p, q := m0.Dimension()
// 	m := m0.Apply(m0.Dagger())
// 	e := Eps(eps...)

// 	for i := 0; i < p; i++ {
// 		for j := 0; j < q; j++ {
// 			if i == j {
// 				if cmplx.Abs(m[i][j]-complex(1, 0)) > e {
// 					return false
// 				}
// 				continue
// 			}

// 			if cmplx.Abs(m[i][j]-complex(0, 0)) > e {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func (m0 Matrix) Apply(m1 Matrix) Matrix {
// 	m, n := m1.Dimension()
// 	p, _ := m0.Dimension()

// 	m2 := Matrix{}
// 	for i := 0; i < m; i++ {
// 		v := []complex128{}
// 		for j := 0; j < n; j++ {
// 			c := complex(0, 0)
// 			for k := 0; k < p; k++ {
// 				c = c + m1[i][k]*m0[k][j]
// 			}
// 			v = append(v, c)
// 		}
// 		m2 = append(m2, v)
// 	}

// 	return m2
// }

// func (m0 Matrix) Mul(z complex128) Matrix {
// 	p, q := m0.Dimension()

// 	m := Matrix{}
// 	for i := 0; i < p; i++ {
// 		v := []complex128{}
// 		for j := 0; j < q; j++ {
// 			v = append(v, z*m0[i][j])
// 		}
// 		m = append(m, v)
// 	}

// 	return m
// }

// func (m0 Matrix) Add(m1 Matrix) Matrix {
// 	p, q := m0.Dimension()

// 	m := Matrix{}
// 	for i := 0; i < p; i++ {
// 		v := []complex128{}
// 		for j := 0; j < q; j++ {
// 			v = append(v, m0[i][j]+m1[i][j])
// 		}
// 		m = append(m, v)
// 	}

// 	return m
// }

// func (m0 Matrix) Sub(m1 Matrix) Matrix {
// 	p, q := m0.Dimension()

// 	m := Matrix{}
// 	for i := 0; i < p; i++ {
// 		v := []complex128{}
// 		for j := 0; j < q; j++ {
// 			v = append(v, m0[i][j]-m1[i][j])
// 		}
// 		m = append(m, v)
// 	}

// 	return m
// }

// func (m0 Matrix) Trace() complex128 {
// 	p, _ := m0.Dimension()
// 	var sum complex128
// 	for i := 0; i < p; i++ {
// 		sum = sum + m0[i][i]
// 	}
// 	return sum
// }

// func (m0 Matrix) TensorProduct(m1 Matrix) Matrix {
// 	m, n := m0.Dimension()
// 	p, q := m1.Dimension()

// 	tmp := []Matrix{}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			tmp = append(tmp, m1.Mul(m0[i][j]))
// 		}
// 	}

// 	m2 := Matrix{}
// 	for l := 0; l < len(tmp); l = l + m {
// 		for j := 0; j < p; j++ {
// 			v := []complex128{}
// 			for i := l; i < l+m; i++ {
// 				for k := 0; k < q; k++ {
// 					v = append(v, tmp[i][j][k])
// 				}
// 			}
// 			m2 = append(m2, v)
// 		}
// 	}

// 	return m2
// }

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

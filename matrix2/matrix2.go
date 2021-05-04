package matrix2

import (
	"math"
)

// Matrix2 represents a 2x2 Matrix of float64 with matrix entries in the form
// Mij with i(row),j(column) varying from 1 to 2
type Matrix2 struct {
	M11, M12, M21, M22 float64
}

// NewIdentity returns a Matrix2 identity matrix
func NewIdentity() *Matrix2 {
	return &Matrix2{1, 0, 0, 1}
}

// New is a variadic function that takes 4 args: M11, M12, M21, M22 and returns
// a new Matrix2 or an error
func New(m11, m12, m21, m22 float64) *Matrix2 {
	return &Matrix2{m11, m12, m21, m22}
}

// ToArray takes a Matrix2 and returns a 2x2 Array of Arrays
func (m2 Matrix2) ToArray() [2][2]float64 {
	return [2][2]float64{{m2.M11, m2.M12}, {m2.M21, m2.M22}}
}

// ToFlatArray takes a Matrix2 and returns an Array with length 4 => [m11, m12, m21, m22]
func (m2 Matrix2) ToFlatArray() [4]float64 {
	return [4]float64{m2.M11, m2.M12, m2.M21, m2.M22}
}

// RotationMatrix takes an angle in radians and returns a rotation matrix.
// Note: positive angles generate a counter-clockwise rotation
//  		 negative angles generate a clockwise rotation
func RotationMatrix(theta float64) *Matrix2 {
	return &Matrix2{
		math.Cos(theta),
		-math.Sin(theta),
		math.Sin(theta),
		math.Cos(theta),
	}
}

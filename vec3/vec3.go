package vec3

import (
	"math"
)

// Vec3 defines a 3-dimension vector
type Vec3 struct {
	I float64
	J float64
	K float64
}

// New returns a new Vec3 pointer
func New(i float64, j float64, k float64) *Vec3 {
	return &Vec3{i, j, k}
}

// Add takes a vec3 transforming v1 by adding their dimensions
func (v1 *Vec3) Add(v2 Vec3) {
	v1.I += v2.I
	v1.J += v2.J
	v1.K += v2.K
}

// Add takes two vec3 returning a new vec3 whose dimensions are equal to the sum of the given vec3 dimensions
func Add(v1 Vec3, v2 Vec3) *Vec3 {
	return &Vec3{v1.I + v2.I, v1.J + v2.J, v1.K + v2.K}
}

// Dot takes a vec3 and returns the dot product of the two vectors
func (v1 *Vec3) Dot(v2 Vec3) float64 {
	return Dot(*v1, v2)
}

// Dot takes two vec3 and returns the dot product of the two vec3
func Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.I*v2.I + v1.J*v2.J + v1.K*v2.K
}

// Magnitude returns the magnitude of a vec3
func (v1 *Vec3) Magnitude() float64 {
	return Magnitude(*v1)
}

// Magnitude takes a vec3 and returns its magnitude
func Magnitude(v Vec3) float64 {
	return math.Sqrt(Dot(v, v))
}

// MirrorI transforms v by mirroring its I-Dimension
func (v1 *Vec3) MirrorI() {
	v1.I = -v1.I
}

// MirrorI takes a Vec3 and mirrors its I-dimension returning a new Vec3
func MirrorI(v Vec3) *Vec3 {
	return &Vec3{-v.I, v.J, v.K}
}

// MirrorJ transforms v by mirroring its J-Dimension
func (v1 *Vec3) MirrorJ() {
	v1.J = -v1.J
}

// MirrorJ takes a Vec3 and mirrors its J-dimension returning a new Vec3
func MirrorJ(v Vec3) *Vec3 {
	return &Vec3{v.I, -v.J, v.K}
}

// MirrorK transforms v by mirroring its K-Dimension
func (v1 *Vec3) MirrorK() {
	v1.K = -v1.K
}

// MirrorK takes a Vec3 and mirrors its K-dimension returning a new Vec3
func MirrorK(v Vec3) *Vec3 {
	return &Vec3{v.I, v.J, -v.K}
}

// Mirror transforms v by mirroring all of its dimensions
func (v1 *Vec3) Mirror() {
	v1.MirrorI()
	v1.MirrorJ()
	v1.MirrorK()
}

// Mirror takes a Vec3 and mirrors all of its dimensions, returning a new Vec3
func Mirror(v Vec3) *Vec3 {
	return &Vec3{-v.I, -v.J, -v.K}
}

// Scale scales v by an s amount
func (v1 *Vec3) Scale(s float64) {
	v1.I *= s
	v1.J *= s
	v1.K *= s
}

// Scale scales a given Vec3 by an s amount, returning a new Vec3
func Scale(v Vec3, s float64) *Vec3 {
	return &Vec3{v.I * s, v.J * s, v.K * s}
}

// Sum takes a Vec3 and adds it to the underlying Vec3 transforming it
func (v1 *Vec3) Sum(v2 Vec3) {
	v1.I += v2.I
	v1.J += v2.J
	v1.K += v2.K
}

// Sum takes two Vec3 and returns their sum as a new Vec3
func Sum(v1 Vec3, v2 Vec3) *Vec3 {
	return &Vec3{v1.I + v2.I, v1.J + v2.J, v1.K + v2.K}
}

// Unit tranforms a Vec3 into a unit vector
func (v1 *Vec3) Unit() {
	m := Magnitude(*v1)

	v1.I /= m
	v1.J /= m
	v1.K /= m
}

// Unit takes a Vec3 and returns a unit vector
func Unit(v Vec3) *Vec3 {
	m := Magnitude(v)

	return &Vec3{v.I / m, v.J / m, v.K / m}
}

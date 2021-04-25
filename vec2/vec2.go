package vec2

import (
	"math"
)

// Vec2 defines a 3-dimension vector
type Vec2 struct {
	I float64
	J float64
}

// New returns a nre vec2 pointer
func New(i float64, j float64) *Vec2 {
	return &Vec2{i, j}
}

// Add takes a Vec2 transforming v1 by adding their dimensions
func (v1 *Vec2) Add(v2 Vec2) {
	v1.I += v2.I
	v1.J += v2.J
}

// Add takes two Vec2 returning a new Vec2 whose dimensions are equal to the sum of the given Vec2 dimensions
func Add(v1 Vec2, v2 Vec2) *Vec2 {
	return &Vec2{v1.I + v2.I, v1.J + v2.J}
}

// Dot takes a Vec2 and returns the dot product of the two vectors
func (v1 *Vec2) Dot(v2 Vec2) float64 {
	return Dot(*v1, v2)
}

// Dot takes two Vec2 and returns the dot product of the two Vec2
func Dot(v1 Vec2, v2 Vec2) float64 {
	return v1.I*v2.I + v1.J*v2.J
}

// Magnitude returns the magnitude of a Vec2
func (v1 *Vec2) Magnitude() float64 {
	return Magnitude(*v1)
}

// Magnitude takes a Vec2 and returns its magnitude
func Magnitude(v Vec2) float64 {
	return math.Sqrt(Dot(v, v))
}

// MirrorI transforms v by mirroring its I-Dimension
func (v1 *Vec2) MirrorI() {
	v1.I = -v1.I
}

// MirrorI takes a Vec2 and mirrors its I-dimension returning a new Vec2
func MirrorI(v Vec2) *Vec2 {
	return &Vec2{-v.I, v.J}
}

// MirrorJ transforms v by mirroring its J-Dimension
func (v1 *Vec2) MirrorJ() {
	v1.J = -v1.J
}

// MirrorJ takes a Vec2 and mirrors its J-dimension returning a new Vec2
func MirrorJ(v Vec2) *Vec2 {
	return &Vec2{v.I, -v.J}
}

// Mirror transforms v by mirroring all of its dimensions
func (v1 *Vec2) Mirror() {
	v1.MirrorI()
	v1.MirrorJ()
}

// Mirror takes a Vec2 and mirrors all of its dimensions, returning a new Vec2
func Mirror(v Vec2) *Vec2 {
	return &Vec2{-v.I, -v.J}
}

// Scale scales v by an s amount
func (v1 *Vec2) Scale(s float64) {
	v1.I *= s
	v1.J *= s
}

// Scale scales a given Vec2 by an s amount, returning a new Vec2
func Scale(v Vec2, s float64) *Vec2 {
	return &Vec2{v.I * s, v.J * s}
}

// Sum takes a Vec2 and adds it to the underlying Vec2 transforming it
func (v1 *Vec2) Sum(v2 Vec2) {
	v1.I += v2.I
	v1.J += v2.J
}

// Sum takes two Vec2 and returns their sum as a new Vec2
func Sum(v1 Vec2, v2 Vec2) *Vec2 {
	return &Vec2{v1.I + v2.I, v1.J + v2.J}
}

// Unit tranforms a Vec2 into a unit vector
func (v1 *Vec2) Unit() {
	m := Magnitude(*v1)

	v1.I /= m
	v1.J /= m
}

// Unit takes a Vec2 and returns a unit vector
func Unit(v Vec2) *Vec2 {
	m := Magnitude(v)

	return &Vec2{v.I / m, v.J / m}
}

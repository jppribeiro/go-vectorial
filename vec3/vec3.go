package vec3

import (
	"math"
)

type vec3 struct {
	I float64
	J float64
	K float64
}

func New(i float64, j float64, k float64) *vec3 {
	return &vec3{i, j, k}
}

func (v1 *vec3) Add(v2 vec3) {
	v1.I += v2.I
	v1.J += v2.J
	v1.K += v2.K
}

func Add(v1 vec3, v2 vec3) *vec3 {
	return &vec3{v1.I + v2.I, v1.J + v2.J, v1.K + v2.K}
}

func (v1 *vec3) Dot(v2 vec3) float64 {
	return Dot(*v1, v2)
}

func Dot(v1 vec3, v2 vec3) float64 {
	return v1.I*v2.I + v1.J*v2.J + v1.K*v2.K
}

func (v *vec3) Magnitude() float64 {
	return Magnitude(*v)
}

func Magnitude(v vec3) float64 {
	return math.Sqrt(Dot(v, v))
}

func (v *vec3) MirrorI() {
	v.I = -v.I
}

func MirrorI(v vec3) *vec3 {
	return &vec3{-v.I, v.J, v.K}
}

func (v *vec3) MirrorJ() {
	v.J = -v.J
}

func MirrorJ(v vec3) *vec3 {
	return &vec3{v.I, -v.J, v.K}
}

func (v *vec3) MirrorK() {
	v.K = -v.K
}

func MirrorK(v vec3) *vec3 {
	return &vec3{v.I, v.J, -v.K}
}

func (v *vec3) Mirror() {
	v.MirrorI()
	v.MirrorJ()
	v.MirrorK()
}

func Mirror(v vec3) *vec3 {
	return &vec3{-v.I, -v.J, -v.K}
}

func (v *vec3) Scale(s float64) {
	v.I *= s
	v.J *= s
	v.K *= s
}

func Scale(v vec3, s float64) *vec3 {
	return &vec3{v.I * s, v.J * s, v.K * s}
}

func (v1 *vec3) Sum(v2 vec3) {
	v1.I += v2.I
	v1.J += v2.J
	v1.K += v2.K
}

func Sum(v1 vec3, v2 vec3) *vec3 {
	return &vec3{v1.I + v2.I, v1.J + v2.J, v1.K + v2.K}
}

func (v *vec3) Unit() {
	m := v.Magnitude()

	v.I /= m
	v.J /= m
	v.K /= m
}

func Unit(v vec3) *vec3 {
	m := Magnitude(v)

	return &vec3{v.I / m, v.J / m, v.K / m}
}

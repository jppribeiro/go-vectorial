package vec3

import (
	"math"
	"reflect"
	"testing"
)

func Test_Vec3_Add(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	type args struct {
		v2 Vec3
	}

	v := fields{1, 1, 1}

	args1 := args{Vec3{1, 2, 3}}
	args2 := args{Vec3{-1, 2, 3}}
	args3 := args{Vec3{-2, -3, 3}}
	args4 := args{Vec3{0, 0, 0}}
	args5 := args{Vec3{-1, -1, -1}}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec3
	}{
		{"v2: {1, 2, 3}", v, args1, Vec3{2, 3, 4}},
		{"v2: {-1, 2, 3}", v, args2, Vec3{0, 3, 4}},
		{"v2: {-2, -3, 3}", v, args3, Vec3{-1, -2, 4}},
		{"v2: {0, 0, 0}", v, args4, Vec3{1, 1, 1}},
		{"v2: {-2, -3, 3}", v, args5, Vec3{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			v1.Add(tt.args.v2)

			if !reflect.DeepEqual(*v1, tt.want) {
				t.Errorf("Sum() = %v, want %v", v1, tt.want)
			}
		})
	}
}

func Test_Vec3_Dot(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"for v1:{1, 1, 1} and v2:{1, 1, 1} expect 3", fields{1, 1, 1}, args{Vec3{1, 1, 1}}, 3},
		{"for v1:{2, 3, 4} and v2:{-1, -1, 1} expect -1", fields{2, 3, 4}, args{Vec3{-1, -1, 1}}, -1},
		{"for v1:{Pi, 1, 1} and v2:{1, 2, 0} expect 2Pi", fields{math.Pi, 1, 1}, args{Vec3{1, 2, 0}}, math.Pi + 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			if got := v1.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Vec3.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Vec3_Magnitude(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"With v{1, 1, 1} expect sqr(3)", fields{1, 1, 1}, math.Sqrt(3)},
		{"With v{-1, 1, 1} expect sqr(3)", fields{-1, 1, 1}, math.Sqrt(3)},
		{"With v{-1, 1, -1} expect sqr(3)", fields{-1, 1, -1}, math.Sqrt(3)},
		{"With v{0, 0, 0} expect sqr(0)", fields{0, 0, 0}, math.Sqrt(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			if got := v.Magnitude(); got != tt.want {
				t.Errorf("Vec3.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Vec3_Mirror(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec3
	}{
		{"v{1, 1, 1} -> v{-1, -1, -1}", fields{1, 1, 1}, Vec3{-1, -1, -1}},
		{"v{0, 0, 0} -> v{0, 0, 0}", fields{0, 0, 0}, Vec3{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			v.Mirror()

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec3.Mirror() = %v, want %v", *v, tt.want)
			}
		})
	}
}

func Test_Vec3_Scale(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	type args struct {
		s float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec3
	}{
		{"v{1, 1, 1}; s = 2 -> v{2, 2, 2}", fields{1, 1, 1}, args{2}, Vec3{2, 2, 2}},
		{"v{0, 0, 0}; s = 2 -> v{0, 0, 0}", fields{0, 0, 0}, args{2}, Vec3{0, 0, 0}},
		{"v{1, 1, 1}; s = 0 -> v{0, 0, 0}", fields{1, 1, 1}, args{0}, Vec3{0, 0, 0}},
		{"v{0.501, 1, 1}; s = 2 -> v{1.002, 2, 2}", fields{0.501, 1, 1}, args{2}, Vec3{1.002, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			v.Scale(tt.args.s)

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec3.Scale() = %v, want %v", *v, tt.want)
			}
		})
	}
}

func Test_Vec3_Sum(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec3
	}{
		{"v1{1, 1, 1} + v2{1, 1, 1} -> v1{2, 2, 2}", fields{1, 1, 1}, args{Vec3{1, 1, 1}}, Vec3{2, 2, 2}},
		{"v1{1, 1, 1} + v2{-1, -1, -1} -> v1{0, 0, 0}", fields{1, 1, 1}, args{Vec3{-1, -1, -1}}, Vec3{0, 0, 0}},
		{"v1{0, 0, 0} + v2{-1, -1, -1} -> v1{-1, -1, -1}", fields{0, 0, 0}, args{Vec3{-1, -1, -1}}, Vec3{-1, -1, -1}},
		{"v1{1, 1, 1} + v2{0, 0, 0} -> v1{1, 1, 1}", fields{1, 1, 1}, args{Vec3{0, 0, 0}}, Vec3{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			v1.Sum(tt.args.v2)

			if !reflect.DeepEqual(*v1, tt.want) {
				t.Errorf("Vec3.Scale() = %v, want %v", *v1, tt.want)
			}
		})
	}
}

func Test_Vec3_Unit(t *testing.T) {
	type fields struct {
		I float64
		J float64
		K float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec3
	}{
		{"v{2,0,0} -> v{1,0,0}", fields{2, 0, 0}, Vec3{1, 0, 0}},
		{"v{1,1,1} -> v{1/sqr(3),1/sqr(3),1/sqr(3)}", fields{1, 1, 1}, Vec3{1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			v.Unit()

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec3.Unit() = %v, want %v", *v, tt.want)
			}
		})
	}
}

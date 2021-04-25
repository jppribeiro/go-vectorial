package vec2

import (
	"math"
	"reflect"
	"testing"
)

func Test_Vec2_Add(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	type args struct {
		v2 Vec2
	}

	v := fields{1, 1}

	args1 := args{Vec2{1, 2}}
	args2 := args{Vec2{-1, 2}}
	args3 := args{Vec2{-2, -3}}
	args4 := args{Vec2{0, 0}}
	args5 := args{Vec2{-1, -1}}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{"v2: {1, 2}", v, args1, Vec2{2, 3}},
		{"v2: {-1, 2}", v, args2, Vec2{0, 3}},
		{"v2: {-2, -3}", v, args3, Vec2{-1, -2}},
		{"v2: {0, 0}", v, args4, Vec2{1, 1}},
		{"v2: {-2, -3}", v, args5, Vec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			v1.Add(tt.args.v2)

			if !reflect.DeepEqual(*v1, tt.want) {
				t.Errorf("Sum() = %v, want %v", v1, tt.want)
			}
		})
	}
}

func Test_Vec2_Dot(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"for v1:{1, 1} and v2:{1, 1} expect 2", fields{1, 1}, args{Vec2{1, 1}}, 2},
		{"for v1:{2, 3} and v2:{-1, -1} expect -5", fields{2, 3}, args{Vec2{-1, -1}}, -5},
		{"for v1:{Pi, 1} and v2:{1, 2} expect 2Pi", fields{math.Pi, 1}, args{Vec2{1, 2}}, math.Pi + 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			if got := v1.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Vec2.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Vec2_Magnitude(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"With v{1, 1} expect sqr(2)", fields{1, 1}, math.Sqrt(2)},
		{"With v{-1, 1} expect sqr(2)", fields{-1, 1}, math.Sqrt(2)},
		{"With v{-1, 1} expect sqr(2)", fields{-1, 1}, math.Sqrt(2)},
		{"With v{0, 0} expect sqr(0)", fields{0, 0}, math.Sqrt(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			if got := v.Magnitude(); got != tt.want {
				t.Errorf("Vec2.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Vec2_Mirror(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec2
	}{
		{"v{1, 1} -> v{-1, -1}", fields{1, 1}, Vec2{-1, -1}},
		{"v{0, 0} -> v{0, 0}", fields{0, 0}, Vec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			v.Mirror()

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec2.Mirror() = %v, want %v", *v, tt.want)
			}
		})
	}
}

func Test_Vec2_Scale(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	type args struct {
		s float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{"v{1, 1}; s = 2 -> v{2, 2}", fields{1, 1}, args{2}, Vec2{2, 2}},
		{"v{0, 0}; s = 2 -> v{0, 0}", fields{0, 0}, args{2}, Vec2{0, 0}},
		{"v{1, 1}; s = 0 -> v{0, 0}", fields{1, 1}, args{0}, Vec2{0, 0}},
		{"v{0.501, 1}; s = 2 -> v{1.002, 2}", fields{0.501, 1}, args{2}, Vec2{1.002, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			v.Scale(tt.args.s)

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec2.Scale() = %v, want %v", *v, tt.want)
			}
		})
	}
}

func Test_Vec2_Sum(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{"v1{1, 1} + v2{1, 1} -> v1{2, 2}", fields{1, 1}, args{Vec2{1, 1}}, Vec2{2, 2}},
		{"v1{1, 1} + v2{-1, -1} -> v1{0, 0}", fields{1, 1}, args{Vec2{-1, -1}}, Vec2{0, 0}},
		{"v1{0, 0} + v2{-1, -1} -> v1{-1, -1}", fields{0, 0}, args{Vec2{-1, -1}}, Vec2{-1, -1}},
		{"v1{1, 1} + v2{0, 0} -> v1{1, 1}", fields{1, 1}, args{Vec2{0, 0}}, Vec2{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			v1.Sum(tt.args.v2)

			if !reflect.DeepEqual(*v1, tt.want) {
				t.Errorf("Vec2.Scale() = %v, want %v", *v1, tt.want)
			}
		})
	}
}

func Test_Vec2_Unit(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec2
	}{
		{"v{2,0,0} -> v{1,0,0}", fields{2, 0}, Vec2{1, 0}},
		{"v{1,1,1} -> v{1/sqr(2),1/sqr(2),1/sqr(2)}", fields{1, 1}, Vec2{1 / math.Sqrt(2), 1 / math.Sqrt(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			v.Unit()

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("Vec2.Unit() = %v, want %v", *v, tt.want)
			}
		})
	}
}

func TestVec2_Angle(t *testing.T) {
	type fields struct {
		I float64
		J float64
	}
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"{1,0}^{0,4} -> expect pi/2", fields{1, 0}, args{Vec2{0, 4}}, math.Pi / 2},
		{"{2,0}^{1,0} -> expect 0", fields{2, 0}, args{Vec2{1, 0}}, 0},
		{"{1,0}^{-2,0} -> expect 0", fields{1, 0}, args{Vec2{-2, 0}}, math.Pi},
		{"{1,0}^{-0.707.., -0.707..} -> expect 3*pi/4", fields{1, 0}, args{Vec2{-math.Sqrt(2) / 2, -math.Sqrt(2) / 2}}, 3 * math.Pi / 4},
		{"{1,0}^{-0.707.., 0.707..} -> expect 3*pi/4", fields{1, 0}, args{Vec2{-math.Sqrt(2) / 2, math.Sqrt(2) / 2}}, 3 * math.Pi / 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := &Vec2{
				I: tt.fields.I,
				J: tt.fields.J,
			}
			if got := v1.Angle(tt.args.v2); got != tt.want {
				t.Errorf("Vec2.Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}

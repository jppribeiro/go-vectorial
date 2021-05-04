package matrix2

import (
	"reflect"
	"testing"
)

func TestMatrix2_ToArray(t *testing.T) {
	type fields struct {
		M11 float64
		M12 float64
		M21 float64
		M22 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   [2][2]float64
	}{
		{"return [[1,2][3,4]]", fields{1, 2, 3, 4}, [2][2]float64{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m2 := Matrix2{
				M11: tt.fields.M11,
				M12: tt.fields.M12,
				M21: tt.fields.M21,
				M22: tt.fields.M22,
			}
			if got := m2.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix2.ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix2_ToFlatArray(t *testing.T) {
	type fields struct {
		M11 float64
		M12 float64
		M21 float64
		M22 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]float64
	}{
		{"return [1, 2, 3, 4]", fields{1, 2, 3, 4}, [4]float64{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m2 := Matrix2{
				M11: tt.fields.M11,
				M12: tt.fields.M12,
				M21: tt.fields.M21,
				M22: tt.fields.M22,
			}
			if got := m2.ToFlatArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix2.ToFlatArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		m11 float64
		m12 float64
		m21 float64
		m22 float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix2
	}{
		{"return Matrix2{1, 2, 3, 4}", args{1, 2, 3, 4}, &Matrix2{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.m11, tt.args.m12, tt.args.m21, tt.args.m22); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIdentity(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix2
	}{
		{"Return 2x2 Identity Matrix", &Matrix2{1, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdentity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

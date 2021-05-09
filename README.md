# Go-Vectorial

[![Go Report Card](https://goreportcard.com/badge/github.com/jppribeiro/go-vectorial)](https://goreportcard.com/report/github.com/jppribeiro/go-vectorial) [![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)

## About <a name = "about"></a>

```go-vectorial``` is a Go package that provides simple 2D & 3D vector algebra. It is intended to be as flexible as possible on its usage.


## Getting Started <a name = "getting_started"></a>

To use one of the packages run:

```go get github.com/jppribeiro/go-vectorial/vec3```
```go get github.com/jppribeiro/go-vectorial/vec2```
```go get github.com/jppribeiro/go-vectorial/matrix2```

## Usage <a name = "usage"></a>

To get the value of a specific coordinate:

```go
v1 := &vec3.Vec3{I: 1, J: 2, K: 3}

j := v1.J

fmt.Println(j) // 2
```

Unit vector
```go
// Create a new 3D vector with coordinates (i,j,k)=(1,1,1)
// The package is implemented with the notation (i,j,k) but you can think in terms of (x,y,z)
v1 := &vec3.Vec3{1,1,1}

fmt.Println(*v) // {1,1,1}

v2 := &vec3.Vec3{1,1,1}

// To calculate a unit vector we have two alternatives

v1.Unit() // transforms v1 into a unit vector

v3 := vec3.Unit(*v2) // Create a new unit vector from an existing vector

fmt.Printf("v1: %v, v2: %v, v3: %v", *v1, *v2, *v3)

// v1: {0.577, 0.577, 0.577}, v2: {1, 1, 1}, v3: {0.577, 0.577, 0.577}

```

Rotate a 2D vector by Pi/2 (90º)

```go

v := &vec.Vec2{2, 0}

v.Rotate(math.Pi/4)

fmt.Printf("%v", *v)

// {0, 2}
```

You can either call the methods on a struct pointer or call the package methods and pass by struct value.


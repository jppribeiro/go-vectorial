# Go-Vectorial

[![Go Report Card](https://goreportcard.com/badge/github.com/jppribeiro/go-vectorial)](https://goreportcard.com/report/github.com/jppribeiro/go-vectorial) [![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)

## About <a name = "about"></a>

```go-vectorial``` is a Go package that provides simple 2D & 3D vector algebra. It is intended to be as flexible as possible on its usage.


## Getting Started <a name = "getting_started"></a>

To use the package run:

```go get github.com/jppribeiro/go-vectorial/vec3```

## Usage <a name = "usage"></a>

```go
import "github.com/jppribeiro/go-vectorial/vec3"

// Create a new 3D vector with coordinates (i,j,k)=(1,1,1)
// The package is implemented with the notation (i,j,k) but you can think in terms of (x,y,z)
v1 := vec3.Vec3{1,1,1}

fmt.Println(*v) // {1,1,1}

v2 := vec3.Vec3{1,1,1}

// To calculate a unit vector we have two alternatives

v1.Unit() // transforms v1 into a unit vector

v3 := vec3.Unit(*v2) // Create a new unit vector from an existing vector

fmt.Printf("v1: %v, v2: %v, v3: %v")

// v1: {0.577, 0.577, 0.577}, v2: {1, 1, 1}, v3: {0.577, 0.577, 0.577}

```

Almost all calculations have these two alternatives.

To get the value of a specific coordinate:

```go
v1 := vec3.Vec3{I: 1, J: 2, K: 3}

j := v1.J

fmt.Println(j) // 2
```
package geometry

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

var zero = Zero()

func Zero() *Vector {
	return &Vector{X: 0, Y: 0}
}

func (v *Vector) Normalize() {
	magnitude := v.Magnitude()
	if magnitude == 0 || magnitude == 1 {
		return
	}
	v.Divide(magnitude)
}

func (v *Vector) Magnitude() float64 {
	return v.DistanceFrom(zero)
}

func (v Vector) String() string {
	return fmt.Sprintf("(%0.0f, %0.0f)", v.X, v.Y)
}

func (v *Vector) AngleTo(other *Vector) float64 {
	dx := other.X - v.X
	dy := other.Y - v.Y
	signedAngle := math.Atan2(dy, dx)
	if signedAngle < 0 {
		signedAngle += 2 * math.Pi
	}
	return signedAngle
}

func (v *Vector) SquareDistanceFrom(other *Vector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return (dx*dx + dy*dy)
}

func (v *Vector) DistanceFrom(other *Vector) float64 {
	return math.Sqrt(v.SquareDistanceFrom(other))
}

func Add(a, b *Vector) *Vector {
	return &Vector{X: a.X + b.X, Y: a.Y + b.Y}
}

func (v *Vector) Add(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector) Subtract(other *Vector) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vector) AddAll(others ...*Vector) {
	for _, other := range others {
		v.Add(other)
	}
}

func (v *Vector) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}

func (v *Vector) Divide(divisor float64) {
	v.X /= divisor
	v.Y /= divisor
}

func (v *Vector) Mod(bounds *Dimension) *Vector {
	v.X = math.Mod(v.X, bounds.W)
	v.Y = math.Mod(v.Y, bounds.H)
	return v
}

func VectorFrom(direction float64, magnitude float64) *Vector {
	sin, cos := math.Sincos(direction)
	return &Vector{X: magnitude * cos, Y: magnitude * sin}
}

func (v *Vector) Limit(max float64) {
	if v.SquareDistanceFrom(zero) > (max * max) {
		v.Normalize()
		v.Scale(max)
	}
}

func Diff(a, b *Vector) *Vector {
	distance := a.DistanceFrom(b)

	result := &Vector{X: a.X, Y: a.Y}
	result.Subtract(b)
	result.Normalize()
	result.Divide(distance)

	return result
}

package forces

import (
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"
)

type Alignment struct {
	BaseForce
}

func (f *Alignment) Init() {
	f.value = geometry.Zero()
	f.count = 0
}

func (f *Alignment) MinDistance() float64 {
	return 300
}

func (a *Alignment) Accumulate(sprite *sprites.Sprite, peer *sprites.Sprite) {
	distance := sprite.Position.DistanceFrom(peer.Position)
	if distance < a.MinDistance() {
		a.count++
		a.value.Add(sprite.Velocity)
	}
}

func (a *Alignment) Finalize(sprite *sprites.Sprite) *geometry.Vector {
	if a.count > 0 {
		a.value.Divide(float64(a.count))
		a.value.Normalize()
		a.value.Scale(MAX_SPEED)
		a.value.Subtract(sprite.Velocity)
		a.value.Limit(MAX_FORCE)
	}
	return a.value
}

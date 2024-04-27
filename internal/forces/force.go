package forces

import (
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"
)

const MAX_SPEED = 2.0
const MAX_FORCE = 0.04

type Force interface {
	Init()
	Accumulate(sprite *sprites.Sprite, peer *sprites.Sprite)
	Finalize(sprite *sprites.Sprite) *geometry.Vector
	MinDistance() float64
}

type BaseForce struct {
	value *geometry.Vector
	count int
}

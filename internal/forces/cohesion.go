package forces

import (
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"
)

type Cohesion struct {
	BaseForce
}

func (c *Cohesion) Init() {
	c.value = geometry.Zero()
	c.count = 0
}

func (c *Cohesion) MinDistance() float64 {
	return 300
}

func (c *Cohesion) Accumulate(sprite *sprites.Sprite, peer *sprites.Sprite) {
	distance := sprite.Position.DistanceFrom(peer.Position)
	if distance < c.MinDistance() {
		c.count++
		c.value.Add(peer.Position)
	}
}

func (c *Cohesion) Finalize(sprite *sprites.Sprite) *geometry.Vector {
	if c.count > 0 {
		c.value.Divide(float64(c.count))
		c.value.Subtract(sprite.Position)
		c.value.Normalize()
		c.value.Scale(MAX_SPEED)
		c.value.Subtract(sprite.Velocity)
		c.value.Limit(MAX_FORCE)
	}
	return c.value
}

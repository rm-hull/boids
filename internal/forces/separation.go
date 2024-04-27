package forces

import (
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"
)

type Separation struct {
	BaseForce
}

func (s *Separation) Init() {
	s.value = geometry.Zero()
	s.count = 0
}

func (s *Separation) MinDistance() float64 {
	return 250
}

func (s *Separation) Accumulate(sprite *sprites.Sprite, peer *sprites.Sprite) {
	distance := sprite.Position.DistanceFrom(peer.Position)
	if distance < s.MinDistance() {
		diff := geometry.Diff(sprite.Position, peer.Position)
		s.value.Add(diff)
		s.count++
	}
}

func (s *Separation) Finalize(sprite *sprites.Sprite) *geometry.Vector {
	if s.count > 0 {
		s.value.Divide(float64(s.count))
	} else if s.value.Magnitude() > 0 {
		s.value.Normalize()
		s.value.Scale(MAX_SPEED)
		s.value.Subtract(sprite.Velocity)
		s.value.Limit(MAX_FORCE)
	} else {
		s.value.Scale(1.5)
	}
	return s.value
}

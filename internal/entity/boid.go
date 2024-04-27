package entity

import (
	"github.com/rm-hull/boids/internal/forces"
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"

	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Boid struct {
	index        int
	flock        []*Boid
	sprite       *sprites.Sprite
	screenBounds *geometry.Dimension
}

func NewFlock(n int, screenBounds *geometry.Dimension) []*Boid {
	var flock = make([]*Boid, n)
	for i := 0; i < n; i++ {
		flock[i] = NewBoid(i, flock, screenBounds)
	}
	return flock
}

func NewBoid(index int, flock []*Boid, screenBounds *geometry.Dimension) *Boid {

	sprite := sprites.NewSprite(screenBounds, sprites.Boid, true)
	sprite.Speed = (rand.Float64() + 0.3) * forces.MAX_SPEED
	sprite.Direction = rand.Float64() * 2 * math.Pi
	sprite.Orientation = sprite.Direction
	sprite.Position.X = rand.Float64() * screenBounds.W
	sprite.Position.Y = rand.Float64() * screenBounds.H
	sprite.Velocity = geometry.VectorFrom(sprite.Direction, sprite.Speed)
	return &Boid{
		index:        index,
		flock:        flock,
		sprite:       sprite,
		screenBounds: screenBounds,
	}
}

func (b *Boid) Draw(screen *ebiten.Image) {
	b.sprite.Draw(screen)
}

func (b *Boid) Update() error {
	b.applyForces(
		&forces.Separation{},
		&forces.Alignment{},
		&forces.Cohesion{},
	)

	if err := b.sprite.Update(); err != nil {
		return err
	}

	b.sprite.Orientation = geometry.Zero().AngleTo(b.sprite.Velocity)
	b.sprite.Velocity.Limit(forces.MAX_SPEED)
	b.sprite.Acceleration.X = 0
	b.sprite.Acceleration.Y = 0

	return nil
}

func (b *Boid) applyForces(forces ...forces.Force) {

	for _, force := range forces {
		force.Init()

		for index, peer := range b.flock {
			if index == b.index {
				continue
			}
			force.Accumulate(b.sprite, peer.sprite)
		}
		value := force.Finalize(b.sprite)

		b.sprite.Acceleration.Add(value)
	}
}

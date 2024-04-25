package entity

import (
	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/internal/sprites"

	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const boidMaxSpeed = 2

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
	sprite.Speed = (rand.Float64() + 0.3) * boidMaxSpeed
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

func (a *Boid) Draw(screen *ebiten.Image) {
	a.sprite.Draw(screen)
}

func (a *Boid) Update() error {
	if err := a.sprite.Update(); err != nil {
		return err
	}
	return nil
}

func (a *Boid) Position() *geometry.Vector {
	return geometry.Add(a.sprite.Position, a.sprite.Centre).Mod(a.screenBounds)
}

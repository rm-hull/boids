package main

import (
	"errors"

	"github.com/rm-hull/boids/internal/entity"
	"github.com/rm-hull/boids/internal/geometry"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	fullscreen bool
	paused     bool
	Flock      []*entity.Boid
}

var screenSize = geometry.Dimension{W: 1024, H: 768}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("dejar de ser un desertor")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.fullscreen = !g.fullscreen
		ebiten.SetFullscreen(g.fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
	}

	if g.paused {
		return nil
	}

	for _, boid := range g.Flock {
		err := boid.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range g.Flock {
		boid.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(screenSize.W), int(screenSize.H)
}

func main() {
	g := &Game{
		fullscreen: false,
		Flock:      entity.NewFlock(30, &screenSize),
	}

	// ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(int(screenSize.W), int(screenSize.H))
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

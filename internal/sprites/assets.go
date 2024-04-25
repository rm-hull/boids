package sprites

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/rm-hull/boids/internal/geometry"
	"github.com/rm-hull/boids/resources/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteSheet = mustLoadImage(images.Asteroids_png)

func mustLoadImage(b []byte) *ebiten.Image {
	r := bytes.NewReader(b)
	img, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func sprite(spriteSheet *ebiten.Image, x, y, w, h int) *ebiten.Image {
	img := spriteSheet.SubImage(image.Rectangle{
		image.Point{x, y},
		image.Point{x + w, y + h},
	})

	return ebiten.NewImageFromImage(img)
}

var Boid = sprite(spriteSheet, 96, 127, 48, 32)


func Centre(sprite *ebiten.Image) geometry.Vector {
	bounds := sprite.Bounds()
	return geometry.Vector{
		X: float64(bounds.Dx()) / 2,
		Y: float64(bounds.Dy()) / 2,
	}
}

func Size(sprite *ebiten.Image) *geometry.Dimension {
	bounds := sprite.Bounds()
	return &geometry.Dimension{
		W: float64(bounds.Dx()),
		H: float64(bounds.Dy()),
	}
}

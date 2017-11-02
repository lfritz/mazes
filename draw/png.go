// Package draw contains functions to draw mazes.
package draw

import (
	"github.com/lfritz/mazes/grid"
	"image"
	"image/color"
	"image/png"
	"os"
)

// ToImage draws m to a grayscale raster image.
func ToImage(m *grid.Maze) image.Image {
	wallColor := color.Gray{0xd0}
	cellWidth := 6

	// create Image object
	w, h := m.Width(), m.Height()
	widthInPixels := w*cellWidth + 1
	heightInPixels := h*cellWidth + 1
	image := image.NewGray(image.Rect(0, 0, widthInPixels, heightInPixels))

	// draw horizontal walls
	for y := 0; y <= m.Height(); y++ {
		j := y * cellWidth
		for x := 0; x < m.Width(); x++ {
			if *m.WallAbove(x, y) {
				i := x * cellWidth
				for k := 0; k <= cellWidth; k++ {
					image.Set(i+k, j, wallColor)
				}
			}
		}
	}

	// draw vertical walls
	for x := 0; x <= m.Width(); x++ {
		i := x * cellWidth
		for y := 0; y < m.Height(); y++ {
			if *m.WallLeftOf(x, y) {
				j := y * cellWidth
				for k := 0; k <= cellWidth; k++ {
					image.Set(i, j+k, wallColor)
				}
			}
		}
	}

	return image
}

func storeImage(image image.Image, path string) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}()
	png.Encode(w, image)
}

// ToPNG draws m and saves it to a PNG file.
func ToPNG(m *grid.Maze, path string) {
	storeImage(ToImage(m), path)
}

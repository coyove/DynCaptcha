package DynCaptcha

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"time"
)

type traveller struct {
	X, Y             float64
	Radius           float64
	TargetX, TargetY float64
	Target           int
	Hotspot          bool
	FakeDirection    float64
}

func (p *traveller) Travel(per float64, clr color.RGBA) (traveller, color.RGBA) {
	const th float64 = 0.2
	fakeX := p.X + math.Cos(p.FakeDirection)*p.Radius*2
	fakeY := p.Y + math.Sin(p.FakeDirection)*p.Radius*2

	var x, y float64
	if per >= th {
		x = (p.TargetX-fakeX)*(per-th)/(1-th) + fakeX
		y = (p.TargetY-fakeY)*(per-th)/(1-th) + fakeY
	} else {
		x = (fakeX-p.X)*(per/th) + p.X
		y = (fakeY-p.Y)*(per/th) + p.Y
	}

	return traveller{X: x, Y: y}, clr
}

func drawNumber(img *image.Paletted, _x float64, _y float64, num int, clr color.RGBA) {
	type _pos struct {
		X, Y int
	}
	var _draw = func(x, y int, n int) {
		var _buf []_pos
		switch n {
		case 0:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				_pos{-2, 2} /*                                */, _pos{2, 2},
				_pos{-2, 1} /*                                */, _pos{2, 1},
				_pos{-2, 0} /*                                */, _pos{2, 0},
				_pos{-2, -1} /*                               */, _pos{2, -1},
				_pos{-2, -2} /*                               */, _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 1:
			_buf = append(_buf,
				_pos{0, 3},
				_pos{0, 2},
				_pos{0, 1},
				_pos{0, 0},
				_pos{0, -1},
				_pos{0, -2},
				_pos{0, -3},
			)
		case 2:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				/*                                             */ _pos{2, 2},
				/*                                             */ _pos{2, 1},
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				_pos{-2, -1}, /*                                           */
				_pos{-2, -2}, /*                                           */
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 3:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				/*                                             */ _pos{2, 2},
				/*                                             */ _pos{2, 1},
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				/*                                             */ _pos{2, -1},
				/*                                             */ _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 4:
			_buf = append(_buf,
				_pos{-2, 3} /*                                */, _pos{2, 3},
				_pos{-2, 2} /*                                */, _pos{2, 2},
				_pos{-2, 1} /*                                */, _pos{2, 1},
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				/*                                             */ _pos{2, -1},
				/*                                             */ _pos{2, -2},
				/*                                             */ _pos{2, -3},
			)
		case 5:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				_pos{-2, 2}, /*                                             */
				_pos{-2, 1}, /*                                             */
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				/*                                            */ _pos{2, -1},
				/*                                            */ _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 6:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				_pos{-2, 2}, /*                                             */
				_pos{-2, 1}, /*                                             */
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				_pos{-2, -1} /*                               */, _pos{2, -1},
				_pos{-2, -2} /*                               */, _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 7:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				/*                                             */ _pos{2, 2},
				/*                                             */ _pos{2, 1},
				/*                                             */ _pos{2, 0},
				/*                                             */ _pos{2, -1},
				/*                                             */ _pos{2, -2},
				/*                                             */ _pos{2, -3},
			)
		case 8:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				_pos{-2, 2} /*                                */, _pos{2, 2},
				_pos{-2, 1} /*                                */, _pos{2, 1},
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				_pos{-2, -1} /*                               */, _pos{2, -1},
				_pos{-2, -2} /*                               */, _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		case 9:
			_buf = append(_buf,
				_pos{-2, 3}, _pos{-1, 3}, _pos{0, 3}, _pos{1, 3}, _pos{2, 3},
				_pos{-2, 2} /*                                */, _pos{2, 2},
				_pos{-2, 1} /*                                */, _pos{2, 1},
				_pos{-2, 0}, _pos{-1, 0}, _pos{0, 0}, _pos{1, 0}, _pos{2, 0},
				/*                                             */ _pos{2, -1},
				/*                                             */ _pos{2, -2},
				_pos{-2, -3}, _pos{-1, -3}, _pos{0, -3}, _pos{1, -3}, _pos{2, -3},
			)
		}

		for _, p := range _buf {
			img.Set(p.X+x, -p.Y+y, clr)
		}
	}

	if num < 10 {
		_draw(int(_x), int(_y), num)
	} else {
		n1 := int(num / 10)
		n2 := num - n1*10

		_draw(int(_x)-3, int(_y), n1)
		_draw(int(_x)+3, int(_y), n2)
	}
}

func drawCircle(img *image.Paletted, _x float64, _y float64, r float64, clr color.RGBA) {
	for x := _x - r; x <= _x+r; x++ {
		for y := _y - r; y <= _y+r; y++ {
			d := math.Sqrt((x-_x)*(x-_x) + (y-_y)*(y-_y))
			if d <= r {
				img.Set(int(x), int(y), clr)
			}
		}
	}
}

func New(seed int) ([]byte, int) {
	var w, h int = 128, 128
	var size = 7
	var group []traveller
	if seed == 0 {
		seed = int(time.Now().UnixNano())
	}
	_rand := rand.New(rand.NewSource(int64(seed)))

	shuffle := _rand.Perm(size * size)
	hotspot := (_rand.Intn(size-2)+1)*size + _rand.Intn(size-2) + 1

	delta := w / (size + 1)

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			p := traveller{X: float64(j * delta), Y: float64(i * delta)}
			group = append(group, p)
		}
	}

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var whiteColor, redColor = color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 0, 255}

	var images []*image.Paletted
	var delays []int

	imgBlink1 := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	imgBlink2 := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	imgFinal := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	imgNum := image.NewPaletted(image.Rect(0, 0, w, h), palette)

	for i, _ := range group {
		group[i].Target = shuffle[i]
		group[i].TargetX = group[shuffle[i]].X
		group[i].TargetY = group[shuffle[i]].Y
		group[i].FakeDirection = float64(_rand.Intn(12)) * math.Pi * 2 / 12
		group[i].Radius = float64(_rand.Intn(2)) + 3

		g := group[i]

		drawCircle(imgBlink1, g.X, g.Y, g.Radius, whiteColor)
		drawCircle(imgFinal, g.TargetX, g.TargetY, g.Radius, whiteColor)

		if i == hotspot {
			drawCircle(imgBlink2, g.X, g.Y, g.Radius, redColor)
			group[i].Hotspot = true
		} else {
			drawCircle(imgBlink2, g.X, g.Y, g.Radius, whiteColor)
			group[i].Hotspot = false
		}

		drawNumber(imgNum, g.X, g.Y, i+1, whiteColor)
	}

	images = append(images, imgBlink1, imgBlink2, imgBlink1, imgBlink2)
	delays = append(delays, 100, 100, 100, 100)

	steps := 48
	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for _, g := range group {
			tmp, newColor := g.Travel(float64(step)/float64(steps), whiteColor)
			drawCircle(img, tmp.X, tmp.Y, g.Radius, newColor)
		}

		// for x := 0; x < w; x++ {
		// 	for y := 0; y < h; y++ {
		// 		if _rand.Intn(10) == 1 {
		// 			img.Set(x, y, whiteColor)
		// 		}
		// 	}
		// }
		c := 0
		for c < int(w*h/15) {
			x := _rand.Intn(w)
			y := _rand.Intn(h)
			img.Set(x, y, whiteColor)
			c++
		}
	}

	images = append(images, imgFinal, imgNum)
	delays = append(delays, 200, 500)

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	gif.EncodeAll(writer, &gif.GIF{Image: images, Delay: delays})

	return buf.Bytes(), group[hotspot].Target + 1
}

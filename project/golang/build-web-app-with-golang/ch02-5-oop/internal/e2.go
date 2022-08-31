package e

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	Width, Height, Depth float64
	Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.Width * b.Height * b.Depth
}

func (b *Box) SetColor(c Color) {
	b.Color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.0
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.Color
		}
	}

	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func E2print() {
	boxes := BoxList{
		Box{Width: 4, Height: 4, Depth: 4, Color: RED},
		Box{Width: 10, Height: 10, Depth: 1, Color: YELLOW},
		Box{Width: 1, Height: 1, Depth: 20, Color: BLACK},
		Box{Width: 10, Height: 10, Depth: 1, Color: BLUE},
		Box{Width: 10, Height: 30, Depth: 1, Color: WHITE},
		Box{Width: 20, Height: 20, Depth: 20, Color: YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm^3")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].Color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].Color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

package shape

type Rectangle struct {
	Width, Length float32
}

func (rectange *Rectangle) Area() float32 {
	return rectange.Width * rectange.Length
}

func (rectange *Rectangle) Perimeter() float32 {
	return 2 * (rectange.Width + rectange.Length)
}

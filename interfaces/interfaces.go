package interfaces

func Perimeter(rectangle Rectange) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectange) float64 {
	return rectangle.Width * rectangle.Height
}

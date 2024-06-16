package shapes

func PointIntersectRectangle(point Point, rec Rectangle) bool {
	return point.X >= rec.X && point.Y >= rec.Y && point.X <= rec.X+rec.Width && point.Y <= rec.Y+rec.Height
}

func RectangleIntersectRectangle(rec1 Rectangle, rec2 Rectangle) bool {
	return !(rec1.X+rec1.Width < rec2.X || rec1.X > rec2.X+rec1.Width ||
		rec1.Y+rec1.Height < rec2.Y || rec1.Y < rec2.Y+rec2.Height)
}

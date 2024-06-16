package shapes

func (size Size) ToPoint() Point {
	return Point{
		X: size.Width,
		Y: size.Height,
	}
}

func (point Point) ConstMul(n int) Point {
	return Point{
		X: point.X * n,
		Y: point.Y * n,
	}
}

func (point Point) ConstDiv(n int) Point {
	return Point{
		X: point.X / n,
		Y: point.Y / n,
	}
}

func (point Point) ConstAdd(n int) Point {
	return Point{
		X: point.X + n,
		Y: point.Y + n,
	}
}

func (point Point) ConstSub(n int) Point {
	return Point{
		X: point.X - n,
		Y: point.Y - n,
	}
}

func (point Point) PointAdd(point2 Point) Point {
	return Point{
		X: point.X + point2.X,
		Y: point.Y + point2.Y,
	}
}

func (point Point) PointSub(point2 Point) Point {
	return Point{
		X: point.X - point2.X,
		Y: point.Y - point2.Y,
	}
}

func (point Point) PointMul(point2 Point) Point {
	return Point{
		X: point.X * point2.X,
		Y: point.Y * point2.Y,
	}
}

func (point Point) PointDiv(point2 Point) Point {
	return Point{
		X: point.X / point2.X,
		Y: point.Y / point2.Y,
	}
}

func (rectangle Rectangle) Center() Point {
	return Point{
		X: rectangle.X + rectangle.Width/2,
		Y: rectangle.Y + rectangle.Height/2,
	}
}

func (size Size) Center() Point {
	return Point{
		X: size.Width / 2,
		Y: size.Height / 2,
	}
}

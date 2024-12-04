package main

type Coordinate struct {
	X         int
	Y         int
	Direction string
}

func (c *Coordinate) GetUp() Coordinate {
	return Coordinate{c.X, c.Y - 1, "U"}
}

func (c *Coordinate) GetUpRight() Coordinate {
	return Coordinate{c.X + 1, c.Y - 1, "UR"}
}

func (c *Coordinate) GetRight() Coordinate {
	return Coordinate{c.X + 1, c.Y, "R"}
}

func (c *Coordinate) GetDownRight() Coordinate {
	return Coordinate{c.X + 1, c.Y + 1, "DR"}
}

func (c *Coordinate) GetDown() Coordinate {
	return Coordinate{c.X, c.Y + 1, "D"}
}

func (c *Coordinate) GetDownLeft() Coordinate {
	return Coordinate{c.X - 1, c.Y + 1, "DL"}
}

func (c *Coordinate) GetLeft() Coordinate {
	return Coordinate{c.X - 1, c.Y, "L"}
}

func (c *Coordinate) GetUpLeft() Coordinate {
	return Coordinate{c.X - 1, c.Y - 1, "UL"}
}

func (c *Coordinate) Step() Coordinate {
	switch c.Direction {
	case "U":
		return c.GetUp()
	case "UR":
		return c.GetUpRight()
	case "R":
		return c.GetRight()
	case "DR":
		return c.GetDownRight()
	case "D":
		return c.GetDown()
	case "DL":
		return c.GetDownLeft()
	case "L":
		return c.GetLeft()
	case "UL":
		return c.GetUpLeft()
	}
	return Coordinate{}
}

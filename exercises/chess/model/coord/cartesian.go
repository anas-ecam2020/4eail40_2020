package coord

import "fmt"

// Cartesian represent a set of cartesian coordinates, x and y.
type Cartesian struct {
	x, y int
}

// NewCartesian is the constructer for Cartesian.
func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

// Coord returns x if n==0, y if n==1
func (c Cartesian) Coord(n int) (int, error) {
	if n == 0 {
		return c.x, nil
	} else if n == 1 {
		return c.y, nil
	} else {
		return 0, fmt.Errorf("Unknown coord component %d", n)
	}
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c%d", 65+c.x, c.y+1)
}

package map_objects

// Rect is used to define the structure of any
// rooms or spaces in the game map that the player
//  is able to walk through.
type Rect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// newRect creates a new Rect starting at x,y with
// a width of w, and a height of h.
func newRect(x, y, w, h int) *Rect {
	r := new(Rect)
	r.X1 = x
	r.Y1 = y
	r.X2 = x + w
	r.Y2 = y + h
	return r
}

// center returns the coordinates of the center of
// the Rect
func (r *Rect) center() (x, y int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2
	return centerX, centerY
}

// intersect determines whether r and other have any coordinates
// that overlap with each other. A main use of this is to determine
// if our map generation has made two overlapping rooms.
func (r *Rect) intersect(other *Rect) bool {
	return (r.X1 <= other.X2 && r.X2 >= other.X1 &&
		r.Y1 <= other.Y2 && r.Y2 >= other.Y1)
}

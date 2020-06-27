package game

type Tile struct {
	Blocked    bool
	BlockSight bool
	Viewable   bool
	Explored   bool
}

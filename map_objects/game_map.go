package map_objects

import (
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
)

type GameMap struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

func NewGameMap(Width, Height int) *GameMap {
	rand.Seed(time.Now().UnixNano())
	g := GameMap{Width, Height, nil}
	g.initializeTiles()
	return &g
}

func (g *GameMap) initializeTiles() {
	tiles := make([][]*Tile, g.Width)
	for i := range tiles {
		tiles[i] = make([]*Tile, g.Height)
		for j := range tiles[i] {
			tiles[i][j] = &Tile{
				Blocked:    true,
				BlockSight: true,
			}
		}
	}
	g.Tiles = tiles
}

func (g *GameMap) MakeMap(
	maxRooms,
	roomMinSize,
	roomMaxSize,
	mapWidth,
	mapHeight int,
	player *entity.Entity,
	entities *[]*entity.Entity,
	maxMonstersPerRoom int,
) {
	rooms := []*Rect{}
	numRooms := 0

	for r := 0; r < maxRooms; r++ {
		w := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		h := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		// random position without going out of the boundaries of the map
		x := rand.Intn(mapWidth - w - 1)
		y := rand.Intn(mapHeight - h - 1)

		newRoom := newRect(x, y, w, h)
		newRoomIntersects := false
		for _, otherRoom := range rooms {
			if newRoom.intersect(otherRoom) {
				newRoomIntersects = true
				break
			}
		}
		if newRoomIntersects {
			continue
		}
		g.createRoom(newRoom)
		newX, newY := newRoom.center()
		if numRooms == 0 {
			player.X = newX
			player.Y = newY
		} else {
			// all rooms after the first:
			// connect it to the previous room with a tunnel

			// center coordinates of previous room
			prevX, prevY := rooms[numRooms-1].center()
			// create tunnel going horizontal, then vertical, or
			// vertical, then horizontal. Determine which is first
			// based off random value
			if rand.Intn(1) == 1 {
				g.createHTunnel(prevX, newX, prevY)
				g.createVTunnel(prevY, newY, newX)
			} else {
				g.createVTunnel(prevY, newY, newX)
				g.createHTunnel(prevX, newX, prevY)
			}
		}
		g.placeEntities(newRoom, entities, maxMonstersPerRoom)
		rooms = append(rooms, newRoom)
		numRooms++
	}
}

func (g *GameMap) createRoom(room *Rect) {
	for x := room.X1 + 1; x < room.X2; x++ {
		for y := room.Y1 + 1; y < room.Y2; y++ {
			g.Tiles[x][y].Blocked = false
			g.Tiles[x][y].BlockSight = false
		}
	}
}

func (g *GameMap) createHTunnel(x1, x2, y int) {
	minX := x1
	maxX := x2 + 1
	if x2 < minX {
		minX = x2
		maxX = x1 + 1
	}
	for x := minX; x < maxX; x++ {
		g.Tiles[x][y].Blocked = false
		g.Tiles[x][y].BlockSight = false
	}
}

func (g *GameMap) createVTunnel(y1, y2, x int) {
	minY := y1
	maxY := y2 + 1
	if y2 < minY {
		minY = y2
		maxY = y1 + 1
	}
	for y := minY; y < maxY; y++ {
		g.Tiles[x][y].Blocked = false
		g.Tiles[x][y].BlockSight = false
	}
}

func (g *GameMap) placeEntities(room *Rect, entities *[]*entity.Entity, maxMonstersPerRoom int) {
	numOfMonsters := rand.Intn(maxMonstersPerRoom)
	for i := 0; i < numOfMonsters; i++ {
		minX := room.X1 + 1
		maxX := room.X2 - 1
		x := rand.Intn(maxX-minX) + minX
		minY := room.Y1 + 1
		maxY := room.Y2 - 1
		y := rand.Intn(maxY-minY) + minY
		entityInLoc := false
		for j := 0; j < len(*entities); j++ {
			entity := (*entities)[j]
			if entity.X == x && entity.Y == y {
				entityInLoc = true
				break
			}
		}
		if !entityInLoc {
			var monster *entity.Entity
			if rand.Intn(100) < 80 {
				monsterFighter := &entity.Fighter{
					Hp:      10,
					Defense: 0,
					Power:   3,
				}
				monsterAi := &entity.BasicMonster{}
				monster = entity.NewEntity(x, y, "O", rl.DarkGreen, "Orc", true, monsterFighter, monsterAi)
			} else {
				monsterFighter := &entity.Fighter{
					Hp:      16,
					Defense: 1,
					Power:   4,
				}
				monsterAi := &entity.BasicMonster{}
				monster = entity.NewEntity(x, y, "T", rl.DarkGreen, "Troll", true, monsterFighter, monsterAi)
			}
			*entities = append(*entities, monster)
		}
	}
}

func (g *GameMap) IsPlayerBlocked(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return true
	}
	tile := g.Tiles[x][y]
	return tile.Blocked
}

// RecomputeFov sets the Viewable property for each
// tile based on the radius for fov provided.
// Algorithm source:
// http://www.roguebasin.com/index.php?title=Simple_and_accurate_LOS_function_for_BlitzMax
func (g *GameMap) RecomputeFov(centerX, centerY, radius int) {
	// reset all tiles as not viewable
	for x := 0; x < len(g.Tiles); x++ {
		for y := 0; y < len(g.Tiles[x]); y++ {
			tile := g.Tiles[x][y]
			if tile.Viewable {
				tile.Viewable = false
			}
		}
	}
	g.Tiles[centerX][centerY].Viewable = true
	for angle := float64(1); angle <= 360; angle += .18 {
		dist := 0
		// set x, y to + .5 to do calculations from center of tile
		x := float64(centerX) + .5
		y := float64(centerY) + .5
		xMove := math.Cos(angle)
		yMove := math.Sin(angle)

		for {
			x += xMove
			y += yMove
			dist += 1
			if dist > radius || x >= float64(g.Width) || x < 0 || y >= float64(g.Height) || y < 0 {
				break
			}
			tile := g.Tiles[int(x)][int(y)]
			tile.Viewable = true
			if tile.BlockSight {
				break
			}
		}
	}
}

func (g *GameMap) MapIsInFov(x, y int) bool {
	row := g.Tiles[x]
	if row == nil {
		return false
	}
	tile := row[y]
	if tile == nil {
		return false
	}
	return tile.Viewable
}

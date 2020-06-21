package map_objects

import (
	"math/rand"
	"time"

	"github.com/kotojo/roguelike/entity"
)

type GameMap struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

func NewGameMap(width, height int) *GameMap {
	rand.Seed(time.Now().UnixNano())
	g := new(GameMap)
	g.Width = width
	g.Height = height
	g.initializeTiles()
	return g
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

func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight int, player *entity.Entity) {
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

func (g *GameMap) IsPlayerBlocked(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return true
	}
	tile := g.Tiles[x][y]
	return tile.Blocked
}

package game

import (
	"container/heap"
	"container/list"
	"fmt"
)

func AStarPath(gameMap Map, startX, startY, goalX, goalY int) *list.List {
	frontier := make(PriorityQueue, 1)
	start := &Item{
		value:    []int{startX, startY},
		priority: 0,
		index:    0,
	}
	frontier[0] = start
	heap.Init(&frontier)
	cameFrom := make(map[string][]int)
	costSoFar := make(map[string]int)

	for frontier.Len() > 0 {
		item := heap.Pop(&frontier).(*Item)
		current := item.value
		if current[0] == goalX && current[1] == goalY {
			break
		}

		neighbors := gameMap.Neighbors(current[0], current[1])
		for _, neighbor := range neighbors {
			key := fmt.Sprint(neighbor[0], neighbor[1])
			newCost := costSoFar[key] + gameMap.Cost(neighbor[0], neighbor[1], goalX, goalY)
			if costSoFar[key] == 0 || costSoFar[key] > newCost {
				costSoFar[key] = newCost
				priority := newCost
				heap.Push(&frontier, &Item{
					value:    neighbor,
					priority: priority,
				})
				cameFrom[key] = current
			}
		}
	}

	node := cameFrom[fmt.Sprint(goalX, goalY)]
	l := list.New()
	for len(node) > 0 && (node[0] != startX || node[1] != startY) {
		l.PushFront(node)
		node = cameFrom[fmt.Sprint(node[0], node[1])]
	}
	return l
}

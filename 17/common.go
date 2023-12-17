package main

import (
	"bufio"
	"container/heap"
	"os"
)

const (
	left  = -1
	right = 1
	up    = -2
	down  = 2
)

var directions = [4]int{left, right, up, down}

type Edge struct {
	r, c          int
	dir           int
	straightCount int
}

type QueueItem struct {
	cost int
	Edge
	prev *QueueItem
}

type PriorityQueue []QueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(QueueItem)) }
func (pq *PriorityQueue) Pop() any          { i := (*pq)[len(*pq)-1]; *pq = (*pq)[0 : len(*pq)-1]; return i }

func parse() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	var weights [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		weights = append(weights, row)
	}

	return weights
}

func dijkstra(weights [][]int, canGoStraight, canTurnOrStop func(int) bool) int {

	queue := PriorityQueue{
		{weights[0][1], Edge{0, 1, right, 1}, nil},
		{weights[1][0], Edge{1, 0, down, 1}, nil},
	}

	visited := make(map[Edge]interface{})

	for len(queue) > 0 {
		cur := heap.Pop(&queue).(QueueItem)

		if _, ok := visited[cur.Edge]; ok {
			continue
		}
		visited[cur.Edge] = struct{}{}

		for _, d := range directions {
			if d == cur.dir*-1 { // can't go back!
				continue
			} else if d == cur.dir {
				if !canGoStraight(cur.straightCount) {
					continue
				}
			} else {
				if !canTurnOrStop(cur.straightCount) {
					continue
				}
			}

			nei := Edge{dir: d, straightCount: 1}
			if d == cur.dir {
				nei.straightCount = cur.straightCount + 1
			}

			switch d {
			case left:
				nei.r, nei.c = cur.r, cur.c-1
			case right:
				nei.r, nei.c = cur.r, cur.c+1
			case up:
				nei.r, nei.c = cur.r-1, cur.c
			case down:
				nei.r, nei.c = cur.r+1, cur.c
			}

			if _, ok := visited[nei]; !ok &&
				nei.r >= 0 && nei.r < len(weights) && nei.c >= 0 && nei.c < len(weights[nei.r]) {

				neiCost := cur.cost + weights[nei.r][nei.c]

				if nei.r == len(weights)-1 && nei.c == len(weights[nei.r])-1 && canTurnOrStop(nei.straightCount) {
					return cur.cost + weights[nei.r][nei.c]
				}

				heap.Push(&queue, QueueItem{neiCost, nei, &cur})
			}
		}
	}

	return -1
}

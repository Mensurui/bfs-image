package main

import (
	"container/list"
	"fmt"
)

func bfsShortestPath(graph map[rune][]rune, start, target rune) []rune {
	queue := list.New()

	queue.PushBack([]rune{start})

	visited := make(map[rune]bool)
	visited[start] = true

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]rune)
		node := path[len(path)-1] // Get the current node (last element of path)

		if node == target {
			return path
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true

				newPath := append([]rune{}, path...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}

	return []rune{}
}

func main() {
	graph := map[rune][]rune{
		'A': {'B', 'C'},
		'B': {'A', 'D', 'E'},
		'C': {'A', 'F'},
		'D': {'B'},
		'E': {'B', 'F'},
		'F': {'C', 'E'},
	}

	path := bfsShortestPath(graph, 'A', 'E')

	if len(path) > 0 {
		fmt.Printf("Shortest path from A to E: %s\n", string(path))
	} else {
		fmt.Println("No path found from A to E.")
	}
}

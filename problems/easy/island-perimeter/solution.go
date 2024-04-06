package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	count := 0
	for i, r := range grid {
		for j, v := range r {
			if v == 1 {
				count += 4
			}
			if v == 1 && j < len(r)-1 && grid[i][j+1] == 1 { // right land
				count--
			}
			if v == 1 && j > 0 && grid[i][j-1] == 1 { // left land
				count--
			}
			if v == 1 && i > 0 && grid[i-1][j] == 1 { // top land
				count--
			}
			if v == 1 && i < len(grid)-1 && grid[i+1][j] == 1 { // bottom land
				count--
			}
		}
	}
	return count
}

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}) == 16)
	fmt.Println(islandPerimeter([][]int{{1}}) == 4)
	fmt.Println(islandPerimeter([][]int{{1, 0}}) == 4)
}

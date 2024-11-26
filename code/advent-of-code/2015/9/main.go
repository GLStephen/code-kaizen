package main

import "fmt"

type nodeSet struct {
	origins map[string]node
}

type node struct {
	neighbors map[string]int
}

func main() {
	fmt.Println("Hello AOC 2015, Day 9")
}

func solve(input []string) (int, error) {
	nodes, err := parseNodes(input)

	return 0, fmt.Errorf("not implemented")
}

func parseNodes(input []string) (nodeSet, error) {
	for _, line := range input {

	}

	return nodeSet{}, fmt.Errorf("not implemented")

}

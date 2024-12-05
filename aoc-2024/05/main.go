package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var day = 5
var path = "input/data.txt"
// var path = "input/test_data.txt"

func getDataFromFile() (map[int][]int, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	isPages := false
	separator := "|"


	graph := make(map[int][]int)
	updates := [][]int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			isPages = true
			separator = ","
			continue
		}

		line := strings.Split(scanner.Text(), separator)

		data := make([]int, len(line))

		for i := 0; i < len(line); i++ {
			data[i], _ = strconv.Atoi(line[i])
		}

		if isPages {
			updates = append(updates, data)
		} else {
			graph[data[0]] = append(graph[data[0]], data[1])
		}
	}

	return graph, updates
}

// Validate if an update respects the ordering rules
func isValidUpdate(graph map[int][]int, update []int) bool {
	// Build a set of nodes in the update
	inUpdate := make(map[int]bool)
	for _, page := range update {
		inUpdate[page] = true
	}

	// Calculate in-degrees for the subset of the graph
	inDegree := make(map[int]int)
	for _, page := range update {
		inDegree[page] = 0
	}
	for from, neighbors := range graph {
		if !inUpdate[from] {
			continue
		}
		for _, to := range neighbors {
			if inUpdate[to] {
				inDegree[to]++
			}
		}
	}

	// Traverse the update in order and validate dependencies
	for _, page := range update {
		if inDegree[page] > 0 {
			return false // Dependency not satisfied
		}
		// Reduce in-degrees of neighbors
		for _, neighbor := range graph[page] {
			if inUpdate[neighbor] {
				inDegree[neighbor]--
			}
		}
	}
	return true
}

func convertToValidUpdate(graph map[int][]int, update []int) ([]int, error) {
	// Step 1: Build the subgraph for the update
	inUpdate := make(map[int]bool)
	for _, page := range update {
		inUpdate[page] = true
	}

	// Subgraph adjacency list and in-degree map
	subgraph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range update {
		inDegree[page] = 0
	}
	for from, neighbors := range graph {
		if !inUpdate[from] {
			continue
		}
		for _, to := range neighbors {
			if inUpdate[to] {
				subgraph[from] = append(subgraph[from], to)
				inDegree[to]++
			}
		}
	}

	// Step 2: Perform topological sort on the subgraph
	var queue []int
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range subgraph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 3: Check for cycles
	if len(sorted) != len(update) {
		return nil, fmt.Errorf("cycle detected, unable to fully order update")
	}

	return sorted, nil
}

func getMiddleNum(arr []int) int {
	return arr[len(arr)/2]
}

func main() {
	start := time.Now()

	graph, updates := getDataFromFile()

	sumOfValidMiddles := 0
	sumOfInvalidMiddles := 0


	for _, update := range updates {
		if isValidUpdate(graph, update) {
			sumOfValidMiddles += getMiddleNum(update)
		} else {
			validatedUpdate, err := convertToValidUpdate(graph, update)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sumOfInvalidMiddles += getMiddleNum(validatedUpdate)
		}
	}

	elapsed := time.Since(start)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintf(w, "Day\t%d\n", day)
	fmt.Fprintf(w, "Part\t1\t%d\n", sumOfValidMiddles)
	fmt.Fprintf(w, "Part\t2\t%d\n", sumOfInvalidMiddles)
	fmt.Fprintf(w, "Time\t\t%s\n", elapsed)
	w.Flush()
}

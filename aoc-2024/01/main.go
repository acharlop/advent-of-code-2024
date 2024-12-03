package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func getDataColsFromFile() ([]int, []int) {
	path := "input/data.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var left []int
	var right []int
	var num int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		num, _ = strconv.Atoi(line[0])
		left = append(left, num)

		num, _ = strconv.Atoi(line[1])
		right = append(right, num)
	}

	return left, right
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func findDistance(a, b []int) int {
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += abs(a[i], b[i])
	}

	return sum
}

func counts(a []int) map[int]int {
	counts := make(map[int]int, len(a))

	for _, v := range a {
		counts[v]++
	}

	return counts
}

func findSimilarity(a, b []int) int {
	countsB := counts(b)

	sum := 0
	
	for i := 0; i < len(a); i++ {
		sum += a[i] * countsB[a[i]]
	}

	return sum
}

func main() {
	start := time.Now()

	left, right := getDataColsFromFile()

	slices.Sort(left)
	slices.Sort(right)

	distance := findDistance(left, right)

	similarity := findSimilarity(left, right)

	elapsed := time.Since(start)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintf(w, "Day\t1\n")
	fmt.Fprintf(w, "Part\t1\t%d\n", distance)
	fmt.Fprintf(w, "Part\t2\t%d\n", similarity)
	fmt.Fprintf(w, "Time\t\t%s\n", elapsed)
	w.Flush()
}

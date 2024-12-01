package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

func similarity(a, b []int) int {
	countsB := counts(b)

	sum := 0
	
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i], countsB[a[i]], a[i] * countsB[a[i]])
		sum += a[i] * countsB[a[i]]
	}

	return sum
}

func getDataFromFile(path string) ([]int, []int) {
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
		line := strings.Split(scanner.Text(), "   ")

		num, _ = strconv.Atoi(line[0])
		left = append(left, num)

		num, _ = strconv.Atoi(line[1])
		right = append(right, num)
	}

	return left, right
}

func main() {
	left, right := getDataFromFile("data/day1.txt")

	slices.Sort(left)
	slices.Sort(right)

	fmt.Println("Part 1: total distance between lists - ", findDistance(left, right))

	fmt.Println("Part 2: similarity score - ", similarity(left, right))
}

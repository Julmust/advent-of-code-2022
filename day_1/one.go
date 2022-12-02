package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func calcMax(data [][]string) {
	var totals []int
	var output int

	for _, e := range data {
		var sum int
		for _, i := range e {
			val, _ := strconv.Atoi(i)
			sum += val
		}

		totals = append(totals, sum)
	}

	sort.Ints(totals)

	for _, e := range totals[len(totals)-3:] {
		output += e
	}

	fmt.Println(output)
}

func parseLines(data []string) [][]string {
	var output [][]string
	var temp []string

	for _, e := range data {
		if e != "" {
			temp = append(temp, e)
		} else {
			output = append(output, temp)
			temp = nil
		}
	}
	output = append(output, temp)

	return output
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	data := readLines("input_one.txt")
	parsedData := parseLines(data)
	calcMax(parsedData)
}

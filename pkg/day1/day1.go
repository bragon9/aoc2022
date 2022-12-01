package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func Part1() (int, error) {
	input := "./inputs/day1/1.txt"
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("unable to open file %v", input)
	}
	defer file.Close()

	max := 0
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			total = 0
			continue
		}

		amt, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		total += amt
		if total > max {
			max = total
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return max, nil
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Part2() (int, error) {
	input := "./inputs/day1/1.txt"
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("unable to open file %v", input)
	}
	defer file.Close()

	minHeap := IntHeap{}
	heap.Init(&minHeap)
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			total = 0
			continue
		}

		amt, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		total += amt

		if len(minHeap) < 3 {
			heap.Push(&minHeap, total)
			continue
		}

		if total > minHeap[0] {
			heap.Pop(&minHeap)
			heap.Push(&minHeap, total)
		}

	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	sum := func() int {
		total := 0
		for _, amt := range minHeap {
			total += amt
		}
		return total
	}()

	return sum, nil
}

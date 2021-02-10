package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Shuffle numbers to ensyre random tree each time
	numbers := []int{3, 4, 5, 7, 8, 9, 10}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	// Build tree
	tree := BinaryTree{data: numbers[0]}

	for _, i := range numbers[1:] {
		tree.Push(i)
	}

	// Print numbers from largest to smallest
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, "th largest item is:", tree.FindKthLargest(i))
	}
}

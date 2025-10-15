package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Random Spark Subject:")
	fmt.Println(Subject(rand.Intn(int(subjectCount) - 1)))
}

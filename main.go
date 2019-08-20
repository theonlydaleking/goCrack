package main

import (
	"fmt"
)

var arr = [...]string{"A", "B", "C"}
var maxDepth int = 10
var currentLevel int = 1

func printer(parentString string, myLevel int, currentLevel int) {
	for _, element := range arr {
		if myLevel == currentLevel {
			fmt.Println(parentString + element)
		} else {
			printer(parentString+element, myLevel+1, currentLevel)
		}

	}
}

func main() {
	for currentLevel <= maxDepth {
		printer("", 1, currentLevel)
		currentLevel++
	}
}

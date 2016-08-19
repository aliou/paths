package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, path := range Paths() {
		fmt.Println(path)
	}
}

func Paths() []string {
	dirs := strings.Split(os.Getenv("PATH"), ":")

	return dirs
}

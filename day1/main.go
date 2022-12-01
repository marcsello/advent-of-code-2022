package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var data string

func main() {

	var acc int = 0
	var top [3]int
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			if acc > top[0] {
				top[2] = top[1]
				top[1] = top[0]
				top[0] = acc
			}
			acc = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		acc += val
	}
	fmt.Println(top[0])
	fmt.Println(top[1])
	fmt.Println(top[2])
	fmt.Println(top[0] + top[1] + top[2])
}

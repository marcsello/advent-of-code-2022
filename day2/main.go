package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	A,X = Rock
	B,Y = Paper
	C,Z = Scissors
*/

func calcScore(opponent, you byte) int {
	score_for_choice := int(you - 'W')

	win_score := 0
	if opponent == (you - 23) {
		win_score = 3
	}

	win_conditions := [][]byte{
		{'A', 'Y'},
		{'B', 'Z'},
		{'C', 'X'},
	}

	for _, match := range win_conditions {
		if (opponent == match[0]) && (you == match[1]) {
			win_score = 6
		}
	}

	return win_score + score_for_choice
}

/*
	X = loose
	Y = draw
	Z = win
*/

func translateChoice(opponent, outcome byte) byte {

	if outcome == 'Y' {
		return opponent + 23
	}

	win_conditions := map[byte]byte {
		'A': 'Y',
		'B': 'Z',
		'C': 'X',
	}

	loose_conditions := map[byte]byte {
		'A': 'Z',
		'B': 'X',
		'C': 'Y',
	}

	if outcome == 'Z' {
		return win_conditions[opponent]
	}

	if outcome == 'X' {
		return loose_conditions[opponent]
	}

	panic("wtf")

}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	scorev1 := 0
	scorev2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		a := line[0]
		b := line[2]
		scorev1 += calcScore(a, b)
		scorev2 += calcScore(a, translateChoice(a, b))
	}

	readFile.Close()

	fmt.Println(scorev1)
	fmt.Println(scorev2)
}

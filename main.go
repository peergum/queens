package main

import _ "net/http/pprof"

import (
	"flag"
	"fmt"
)

var (
	board   [8][8]int
	queens  [8]int
	counter int
)

func main() {
	flag.Parse()

	for c := range 8 {
		clean(0)
		choose(0, c)
	}
	fmt.Println(counter)

}

func choose(r int, c int) {
	if board[r][c] == 0 {
		setQueen(r, c)
	} else {
		return
	}
	if r < 7 {
		for i := range 8 {
			clean(r + 1)
			choose(r+1, i)
		}
	} else {
		counter++
		display()
		return
	}
}

func clean(q int) {
	for r := range 8 {
		for c := range 8 {
			board[r][c] = 0
		}
	}
	for r := range q {
		setQueen(r, queens[r])
	}
}

func setQueen(r int, c int) {
	board[r][c] = 1
	for i := range 8 {
		if i != c {
			board[r][i] = -1
		}
		if i != r {
			board[i][c] = -1
		}
		if r-i >= 0 && c-i >= 0 && board[r-i][c-i] == 0 {
			board[r-i][c-i] = -1
		}
		if r+i < 8 && c+i < 8 && board[r+i][c+i] == 0 {
			board[r+i][c+i] = -1
		}
		if r-i >= 0 && c+i < 8 && board[r-i][c+i] == 0 {
			board[r-i][c+i] = -1
		}
		if r+i < 8 && c-i >= 0 && board[r+i][c-i] == 0 {
			board[r+i][c-i] = -1
		}
	}
	queens[r] = c
}

func display() {
	var x rune
	fmt.Printf("  ")
	for ic := range 8 {
		fmt.Printf("%d ", ic+1)
	}
	fmt.Println()
	for ir, r := range board {
		fmt.Printf("%d ", ir+1)
		for _, c := range r {
			if c == 1 {
				x = 'o'
			} else if c == 0 {
				x = ' '
			} else {
				x = '·'
			}
			fmt.Printf("%c ", x)
		}
		fmt.Println()
	}
	fmt.Println()
}

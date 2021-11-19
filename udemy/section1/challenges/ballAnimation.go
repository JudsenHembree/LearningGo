package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		width  = 20
		height = 10
	)

	var image rune

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	herex :=0
	herey :=0
	velocityx :=1
	velocityy :=1
	board[herex][herey] = true
	for {
		for y := range board[0] {
			for x := range board {
				image = ' ' 
				if board[x][y] {
					image = '\u2713'
				}
				fmt.Print(string(image))
			}
			fmt.Println()
		}
		board[herex][herey]=false
		if herex ==0{
			velocityx = 1
		}
		if herex == 19{
			velocityx=-1
		}
		if herey ==0{
			velocityy = 1
		}
		if herey == 9{
			velocityy=-1
		}
		herex+=velocityx
		herey+=velocityy
		board[herex][herey]=true
		time.Sleep(1 * time.Second)
	}
}

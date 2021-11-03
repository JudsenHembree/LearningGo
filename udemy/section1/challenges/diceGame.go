// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------
import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)


func main() {

	verbose := false

	fmt.Println("line args %d", len(os.Args))
	if(len(os.Args)>=2){
		//fmt.Println(os.Args[1])
		if(strings.ContainsAny(os.Args[1], "-")){
			if(strings.ContainsAny(os.Args[1], "-")){
				verbose=true
			}

		}

	}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Guess a number 1-6")

	winNumber :=rand.Intn(6)+1
	var number int
	//fmt.Println(winNumber)
	fmt.Scanf("%d", &number)
	if number == winNumber{
		fmt.Println("Wow first try")
	}else{
		for number != winNumber{
		fmt.Println("That is incorrect, guess again.")
		if(verbose){
			fmt.Println("The number was", winNumber)
		}
		fmt.Scanf("%d", &number)
		winNumber=rand.Intn(6)+1
		}
	}
	
	win :=make(map[int]string)

	random := rand.Intn(5)

	win[0]="A modern day hero"
	win[1]="Christ you're hot"
	win[2]="Imagine losing, couldn't be me"
	win[3]="Another one"
	win[4]="... Bro your bullying"

	fmt.Println(win[random])
}
package main

import ("fmt")

func main() {
	speed:=10
	force:=20.55

	speed=int(force*float64(speed))
	fmt.Println(speed)
	speed=10
	speed64:=force*float64(speed)
	fmt.Println(speed64)
}

package main

import "fmt"

func duplicateZeros(arr []int)  {
	var sol []int
	for _, num := range arr{
			if num==0{
				if len(sol)<len(arr)-1{
					sol=append(sol,0,0)
				}else{
					sol=append(sol,0)
				}
			}else{
				sol=append(sol,num)
			}
	}

	for index := range arr{
		arr[index]=sol[index]
	}
}

func main(){
	test := []int{1,0,2,3,0,4,5,0}
	var out []int
	duplicateZeros(test)
	correct:=[]int{1,0,0,2,3,0,0,4}
	fmt.Println(correct)
	fmt.Println(test)
	for index:=range out{
		if out[index]==correct[index]{
			fmt.Println("True")
		}
	}
}

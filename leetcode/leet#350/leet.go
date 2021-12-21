package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
    
	used := make(map[int]bool)
	for index := range nums2{
		used[index]=true
	}

    var sol []int
	var set bool
    for _, num := range nums1 {
		set=true;
		for index, num2 := range nums2{
			if num==num2 && set && used[index]{
				set=false
				sol=append(sol,num)
				used[index]=false
			}
		}
	}
	fmt.Println(sol)
	return sol
}


func main(){
	test1:=[]int{1,2,2,1}
	test2:=[]int{2,2}
	intersect(test1, test2)
}

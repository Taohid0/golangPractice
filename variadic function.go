package main

import "fmt"

func sum(a int,nums ...int){
	fmt.Print(nums)
}

func main(){
	
	sum(1,[]int{1,2,3}...)
}

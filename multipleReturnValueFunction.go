package main

import "fmt"

func vals()(int,int,int){
	return 1,2,3
}

func main(){
	_,_,b :=vals()
	fmt.Print(b)
}

package main

import "fmt"

func main(){
	if 7%2==0{
		fmt.Print("7 is even")
	} else{
		fmt.Println("7 is odd")
	}

	if num:=9; num<0{ //a statement can precede conditionals
		fmt.Println(num, "is negative")
	}else{
		fmt.Println(num, "is positive")
	}
}

package main

import "fmt"

func main(){
	nums := []int{2,3,4}
	sum:=0

	for index,num:= range nums{
		sum +=num
		fmt.Print(index)
	}
	fmt.Println(sum)

	kvs := map[string]string{"a":"string","1":"100"}

	for key,value :=range kvs{
		fmt.Println(key,value)
	}
}

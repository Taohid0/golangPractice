package main

import "fmt"

type stackInt struct{
	arr []int
}

func (s *stackInt) push(value int){
	s.arr = append(s.arr, value)
	fmt.Println(s.arr)
}

func (s *stackInt) top() int{
	length := len(s.arr)
	return s.arr[length-1]
}

func (s *stackInt) pop(){
	length := len(s.arr)
	s.arr = s.arr[:length-1]
}

func main()  {

}
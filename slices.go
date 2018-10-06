package main

import "fmt"

func main(){
	s := make([]string,0)
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	s = append(s,"1","3","4")
	s1 := make([]string,len(s))
	copy(s1,s)

	s[0] = "100"
	s1[0] = "200"

	fmt.Println(s,s1)
	fmt.Println(append(s, "e"))
	fmt.Println(s[0])

	t := []string{"g", "h", "i"}
	t = append(t, "a")
	fmt.Println("dcl:", t)
	b := []int{1, 2, 3, 4, 5}
	b = append(b, 3)
	fmt.Println(b)
}

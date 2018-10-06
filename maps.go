package main

import "fmt"

func main()  {
	m := make(map[string]int)

	m["k1"]= 7
	m["kw"] = 13

	fmt.Print("map:",m)

	v1 := m["1"]
	fmt.Println("v1",v1)

	fmt.Println("len:",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	value, exists :=m["k1"]
	fmt.Println(value,exists)

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map",n)
}

package main

import "fmt"

func main() {

	s :=[]int{0,1,2,3,4,5,6,7,8,9}

	s1 :=s[5:8]
	fmt.Println(s1)
	fmt.Println(s[4:])
	n :=copy(s[4:],s1)

	fmt.Println(n,s)

	s2 :=make([]int,6)
	n =copy(s2,s)
	fmt.Println(n,s2)
	
}

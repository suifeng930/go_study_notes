package main

func mkslice() []int {

	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkMap() map[string]int {

	m := make(map[string]int)
	m["a"] = 1
	return m

}
func main() {

	m := mkMap()
	println(m["a"])

	s := mkslice()
	println(s[0])

}

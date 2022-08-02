package main

func main() {
	a := []int{1, 2, 3, 4}
	var ptr [4]*int

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
		println(ptr[i])
	}
}

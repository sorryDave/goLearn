package main

import "fmt"

// single return func
func plus (a int, b int) int {
	return a + b
}

func plusPlus(a,b,c int) int {
	return a + b + c
}

// multiple return func
func vals()(int,int){
	return 3,7
}
// variadic func (args)
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

// closures
func intSeq() func() int {
	i:=0
	return func() int {
		i += 1
		return i
	}
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {

	// single return func
	res:= plus(1,2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)

	// multiple return func
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	//variadic func
	sum(1,2)
	sum(1,2,3)

	nums:=[]int{1,2,3,4}
	sum(nums...)

	// closures
	nextInt := intSeq()

	
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	// recursion
	fmt.Println(fact(2))
}

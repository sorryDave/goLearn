package main
import "fmt"

func zeroval(ival int) {
ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// struct + pointer
type person struct{
	name string
	age int
}


func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// struct pointer
	fmt.Println( person{"bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(&person{name:"Ann", age: 40})

	s := person{name:"Sean", age 48}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)


}

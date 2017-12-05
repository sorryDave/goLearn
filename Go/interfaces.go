package main

import "fmt"

type I interface {
	f1(name string)
	f2(name string) (error, float32)
	f3()
}

type T int64

func (T) f1(name string) {

	fmt.Println(name)
}

func (T) f2(name string) (error, float32) {

	return nil, 10.2
}

func (T) f3() int64 {

	return 10
}

func main() {

}

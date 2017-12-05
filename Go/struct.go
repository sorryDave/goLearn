package main

import (
	"fmt"
	"math"
)

type Rectangle struct{
	x1, y1, x2, y2 float64
}


type Circle struct{
	x float64
	y float64
	z float64
}

// Methods
func (c *Circle) area() float64{
	return math.Pi * c.z * c.z
}


func(r *Rectangle) area() float64{
		l:= distance(r.x1, r.y1, r.x1, r.y2)
		w:= distance(r.x1, r.y1, r.x2, r.y1)
		return l * w
	}


// functions
func circArea(c Circle) float64 {
	return math.Pi * c.z * c.z
}

// pointer function
func circAreaPtr(c *Circle) float64{
	return math.Pi * c.z * c.z
}


func distance(x1, y1, x2, y2 float64) float64{
		a:= x2 - x1
		b:= y2 - y1
		return math.Sqrt(a*a +b*b)
}


// Embedded types
type Person struct{
	Name string

}

func (p *Person) Talk(){
	fmt.Println("Hi, my name is", p.Name)

}



type Android struct{
	Person
	Model string
}

func main(){
	// struct 
	// var c Circle
	// c:= new(Circle)
	c := Circle{x:10, y: 0, z: 8}

	// accessing field
	c.x = 10
	c.y = 6

	fmt.Println(circArea(c))
	fmt.Println(circAreaPtr(&c))

	//Method
	fmt.Println(c.area())

	r := Rectangle{0,0,10,10}
	fmt.Println(r.area())

	// Embedded Types
	a:= new(Android)
	a.Name = "John"
	a.Model = "BF-273"
	a.Person.Talk()
	a.Talk()
}

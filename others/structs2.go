package others

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func changeAge(p *Person) {
	p.age = 30
}

type Android struct {
	Person
	model string
}

type Square struct {
	side float64
}

type Circle struct {
	raduis float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (c *Circle) area() float64 {
	return 3.14 * c.raduis * c.raduis
}

type Shape interface {
	area() float64
}

func Execute() {

	person := Person{name: "Akwasi", age: 21}
	personPtr := &person
	p := new(Person)
	p.name = "Hi"
	fmt.Println(person.name)
	fmt.Println(person.age)

	person.age = 20
	fmt.Println(person.age)
	fmt.Println("Address of person:", personPtr)
	changeAge(&person)
	fmt.Println(person)

	// methods
	person.sayMyName()

	b := Android{model: "Techno"}
	a := new(Android)

	a.Person.sayMyName()
	b.Person.sayMyName()

}
func totalArea(shapes []Shape) float64 {
	total := 0.00

	for _, s := range shapes {
		total += s.area()
	}

	return total
}
func (p *Person) sayMyName() {
	fmt.Println("My name is ", p.name)
}

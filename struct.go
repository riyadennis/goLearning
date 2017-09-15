package main

import "fmt"

// example represents a type with different fields.
type structExample struct {
	counter int16
	counter1 int16
	pi      float32
	flag    bool
}
type Person struct{
	name string
	age int
}
func (p Person) sayHello(){
	fmt.Printf("Hello %s \n", p.name)
}
func (p Person) getAge() (int){
	return p.age
}
func main() {
	//make is not for struct, we can make only channels, slices and maps
	p := Person{"Riya", 22}
	fmt.Println(p.name)
	p.sayHello()
	fmt.Println(p.getAge())

	// Declare a variable of type example set to its
	// zero value.
	var e1 structExample
        e2 := structExample{} // empty literals does not create a value

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 = structExample{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}


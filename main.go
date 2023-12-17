package main

import (
	"fmt"
	"math"
	"pg3/st"
)

type Person struct {
	name string
	age  uint
}

func changePerson(person **Person) {

	*person = &Person{
		name: "vasya",
		age:  25,
	}
	fmt.Printf("change person 2: %+v\n", person)

}

func main() {

	// person := &Person{
	// 	name: "Dm",
	// 	age:  23,
	// }
	// var p *Person = person
	// fmt.Println(&p)

	// changePerson(&person)
	// fmt.Println(person.name)
	fmt.Println(math.Log2(128 * 2))
	list := []int{1, 3, 4, 5, 9}
	fmt.Println(st.Bianary_search(list, 23))
	fmt.Println(st.Bianary_search(list, 5))

	// fmt.Println("max: ", Max(1, 5))

	// fmt.Println(person.name)
	// changePerson(person)
	// fmt.Println(person.name)

	// var creature string = "shark"
	// var pointer *string = &creature

	// fmt.Println("creature: ", creature)
	// fmt.Println("pointer: ", pointer)
	// fmt.Println("*pointer: ", *pointer)
	// *pointer = "dfg"
	// fmt.Println("*pointer: ", *pointer)
	// fmt.Println("creature: ", creature)
	// list := []int{1, 2, 4, 9}
	// fmt.Println(isContains(3, list))
	// fmt.Println(isContains(2, list))

	// fmt.Println(Sum(list))
	// list := []int{1, 2, 4, 9}

	// fmt.Println(st.Ternary(true, 2, 3))
	// fmt.Println(st.Ternary(false, 2, 3))

	// sum := func(x, y int) int { return x + y }
	// mul := func(x, y int) int { return x * y }

	// res := st.Reduce(list, sum, 0)
	// fmt.Println(res)
	// res = st.Reduce(list, mul, 1)
	// fmt.Println(res)

	// var creature *st.Creature
	// creature = &st.Creature{Species: "shark"}
	// fmt.Printf("1) %+v\n", creature)
	// creature.Reset()
	// fmt.Printf("2) %+v\n", creature)

	// fmt.Printf("1) %+v\n", creature)
	// st.ChangeCreature(creature)
	// fmt.Printf("3) %+v\n", creature)
	// fmt.Println(creature.String())

}

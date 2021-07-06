/*Struct
A struct is a collection of fields.
No inheritance, but can use composition via embedding
Tags can be added to struct fields to describe field
*/
package main

import (
	"fmt"
	"reflect"
)

type Actor struct {
	number    int
	actorName string
	episodes  []string
}

type Animal struct {
	Name   string `required max:"100"`
	origin string
}

func main() {
	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"jharkhand":  20304,
		"Tamil Nadu": 86322,
		"Delhi":      7583648,
	}
	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)
	fmt.Println(statePopulations["Delhi"])
	statePopulations["Bihar"] = 8473627
	delete(statePopulations, "jharkhand")
	fmt.Println(statePopulations)

	_, ok := statePopulations["Bihr"]
	fmt.Println(ok)

	v := Actor{
		number:    3,
		actorName: "Amol",
		episodes: []string{
			"hjuh",
			"nhudhu",
		},
	}
	fmt.Println(v)

	aDoctor := struct{ name string }{name: "Amol"}
	anotherDoctor := &aDoctor
	anotherDoctor.name = "KloudOne"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

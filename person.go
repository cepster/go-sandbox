package main

import "fmt"

//Person data struct
type Person struct {
	ID        int8
	FirstName string
	LastName  string
}

func (a Person) idAsString() string {
	return fmt.Sprintf("%d", a.ID)
}

func getSeed() [4]Person {
	return [...]Person{
		Person{1, "Matt", "Richards"},
		Person{2, "Tony", "Lapakko"},
		Person{3, "Ben", "Wanggaard"},
		Person{4, "Nate", "Weimer"},
	}
}

package main

import "fmt"

func main() {

	// Return IPerson
	student, _ := getPerson("student")
	// Return IPerson
	teacher, _ := getPerson("teacher")

	// Convert IPerson to Struct
	studentStruct := student.(*Student)
	fmt.Println("student struct: ", *&studentStruct)

	fmt.Println(student)
	fmt.Println(teacher)
}

type IPerson interface {
	setName(name string)
	getName() string
}

///

type Person struct {
	name string
}

func (g *Person) setName(name string) {
	g.name = name
}

func (g *Person) getName() string {
	return g.name
}

///

type Student struct {
	Person
}

func newStudent() IPerson {
	return &Student{
		Person: Person{
			name: "student 1",
		},
	}
}

///

type Teacher struct {
	Person
}

func newTeacher() IPerson {
	return &Teacher{
		Person: Person{
			name: "teacher 1",
		},
	}
}

///

func getPerson(personType string) (IPerson, error) {
	if personType == "student" {
		return newStudent(), nil
	}
	if personType == "teacher" {
		return newTeacher(), nil
	}
	return nil, fmt.Errorf("Wrong person type passed")
}

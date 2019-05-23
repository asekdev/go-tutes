package main

import (
	"fmt"
)

//Global to the package
var (
	name string = "Andy"
	age  int
)

func main() {

	course := "Java Books"

	fmt.Println("Currently viewing", course)

	//pointing to the memory location of the course variable
	changeCourse(&course)

	fmt.Println("Now viewing", course)

}

//derefencing the paramter, changing its value and returning it
func changeCourse(course *string) string {
	*course = "First Look:  Docker"

	fmt.Println("Attempting to change course to:", *course)

	return *course
}

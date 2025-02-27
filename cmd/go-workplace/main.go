package main

import (
	"fmt"
	"github.com/bluezy47/go-workplace/internal"
)

func main() {
	// Create an instance of EmployeeScoreCalculator
	esc := internal.EmployeeScoreCalculator{}

	// Add hooks
	esc.AddHook(internal.NewDepartmentScoreHook("Engineering"))
	esc.AddHook(internal.NewRoleScoreHook("Manager"))

	// Create an employee
	employee := internal.Employee{
		Name:       "John Doe",
		Age:        30,
		Department: "Engineering",
		Score:      0,
	}
	//
	// Calculate the score
	score := esc.CalculateScore(employee)
	fmt.Println("Final score:", score)
}

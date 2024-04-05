package EmployeeReportWithAi

import (
	"reflect"
	"testing"
)

// prompt 1: show me how to create an array of employees in golan in a single instruction. An employee has an age and a name

type Employee struct {
	Name string
	Age  int
}

var employees = []Employee{
	{Name: "John", Age: 18},
	{Name: "Alice", Age: 30},
	{Name: "Bob", Age: 40},
	{Name: "Mark", Age: 17},
}

// prompt 2: could you make the selected test to pass ?
func TestListAllEmployeesOlderThan18(t *testing.T) {

	var olderEmployees []Employee = ListAllEmployeesOlderThan(employees, 18)

	expectedEmployees := []Employee{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 40},
	}

	if !reflect.DeepEqual(olderEmployees, expectedEmployees) {
		t.Fatalf("expected: %v, got: %v", expectedEmployees, olderEmployees)
	}

}

// prompt 3: could you make the selected test to pass ?
func TestListAllEmployeesYoungerThanAge(t *testing.T) {

	var employeesYoungerThanAge []Employee = ListAllEmployeesYoungerThanAge(employees, 30)

	expectedEmployees := []Employee{
		{Name: "John", Age: 18},
		{Name: "Mark", Age: 17},
	}

	if !reflect.DeepEqual(employeesYoungerThanAge, expectedEmployees) {
		t.Fatalf("expected: %v, got: %v", expectedEmployees, employeesYoungerThanAge)
	}

}

func ListAllEmployeesYoungerThanAge(employees []Employee, age int) []Employee {
	var youngerThanAge []Employee
	for _, emp := range employees {
		if emp.Age < age {
			youngerThanAge = append(youngerThanAge, emp)
		}
	}
	return youngerThanAge
}

func ListAllEmployeesOlderThan(employees []Employee, age int) []Employee {
	var olderEmployees []Employee
	for _, employee := range employees {
		if employee.Age > age {
			olderEmployees = append(olderEmployees, employee)
		}
	}
	return olderEmployees
}

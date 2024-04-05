package EmployeeReportWithAi

import (
	"reflect"
	"testing"
)

// prompt 1: show me how to create an array of employees in golan in a single instruction. An employee has an age and a name

type Employee struct {
	name string
	age  int
}

var employees = []Employee{
	{name: "John", age: 18},
	{name: "Alice", age: 30},
	{name: "Bob", age: 40},
	{name: "Mark", age: 17},
}

// prompt 2: could you make the selected test to pass ?
func TestListAllEmployeesOlderThan(t *testing.T) {

	var olderEmployees []Employee = ListAllEmployeesOlderThan(employees, 18)

	expectedEmployees := []Employee{
		{name: "Alice", age: 30},
		{name: "Bob", age: 40},
	}

	if !reflect.DeepEqual(olderEmployees, expectedEmployees) {
		t.Fatalf("expected: %v, got: %v", expectedEmployees, olderEmployees)
	}

}

// prompt 3: could you make the selected test to pass ?
func TestListAllEmployeesYoungerThanAge(t *testing.T) {

	var employeesYoungerThanAge []Employee = ListAllEmployeesYoungerThan(employees, 30)

	expectedEmployees := []Employee{
		{name: "John", age: 18},
		{name: "Mark", age: 17},
	}

	if !reflect.DeepEqual(employeesYoungerThanAge, expectedEmployees) {
		t.Fatalf("expected: %v, got: %v", expectedEmployees, employeesYoungerThanAge)
	}

}

// prompt 3: is possible eliminate the duplicated code in the selected code ?
func ListAllEmployeesYoungerThan(employees []Employee, ageLimit int) []Employee {
	return filterEmployees(func(employee Employee) bool { return employee.age < ageLimit }, employees)
}

func ListAllEmployeesOlderThan(employees []Employee, ageLimit int) []Employee {
	return filterEmployees(func(employee Employee) bool { return employee.age > ageLimit }, employees)
}

func filterEmployees(condition func(Employee) bool, employees []Employee) []Employee {
	var filteredEmployees []Employee
	for _, emp := range employees {
		if condition(emp) {
			filteredEmployees = append(filteredEmployees, emp)
		}
	}
	return filteredEmployees
}

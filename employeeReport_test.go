package EmployeeReportWithAi

import (
	"reflect"
	"testing"
)

// show me how to create an array of employees in golan in a single instruction. An employee has an age and a name

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

func ListAllEmployeesOlderThan(allEmployees []Employee, ageFilter int) []Employee {
	return []Employee{}
}

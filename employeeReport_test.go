package EmployeeReportWithAi

import (
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

}

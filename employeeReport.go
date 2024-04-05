package TddWithAi

import "github.com/samber/lo"

type Employee struct {
	name string
	age  int
}

// prompt 3: is possible eliminate the duplicated code in the selected code ?
// prompt 4: is possible to refactor the selected code using the filter function in https://github.com/samber/lo?tab=readme-ov-file#filter ?
func ListAllEmployeesYoungerThan(employees []Employee, ageLimit int) []Employee {
	return lo.Filter(employees, func(employee Employee, index int) bool { return employee.age < ageLimit })
}

func ListAllEmployeesOlderThan(employees []Employee, ageLimit int) []Employee {
	return lo.Filter(employees, func(employee Employee, index int) bool { return employee.age > ageLimit })
}

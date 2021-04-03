package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	EmpId    rune   `json:"empId"`
	Role     string `json:"role"`
}

func main() {
	var empId rune

	fmt.Print("Enter id of the employee to search :- ")
	fmt.Scanf("%d", &empId)

	c, err := os.Open("emp.json")
	if err != nil {
		panic(err)
	}

	var emp []Employee
	d := json.NewDecoder(c)
	if err = d.Decode(&emp); err != nil {
		panic(err)
	}

	e, found := searchEmployee(emp, empId)
	if found == false {
		fmt.Println("Employee not found.")
		os.Exit(1)
	}

	e.print()
}

func searchEmployee(employees []Employee, searchId rune) (Employee, bool) {
	found := false
	var emp Employee
	for _, e := range employees {
		if e.EmpId == searchId {
			found = true
			emp = e
			break
		}
	}

	return emp, found
}

func (e Employee) print() {
	fmt.Println("Id: ", e.EmpId)
	fmt.Println("Name: ", e.Name)
	fmt.Println("Location: ", e.Location)
	fmt.Println("Role: ", e.Role)
}

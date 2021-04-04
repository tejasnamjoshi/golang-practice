package main

import (
	"encoding/json"
	"flag"
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
	var f string
	flag.StringVar(&f, "f", "emp.json", "Name of the file with employee records")
	flag.Parse()

	c, err := os.Open(f)
	if err != nil {
		fmt.Println("File not found.")
		os.Exit(1)
		// panic(err)
	}
	defer c.Close()

	fmt.Print("Enter id of the employee to search :- ")
	fmt.Scanf("%d", &empId)

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

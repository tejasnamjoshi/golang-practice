package main

import "fmt"

type Freelance struct {
	empId int
	rate  rune
	days  rune
}

type Permanent struct {
	empId    int
	basicPay rune
	pf       rune
}

type Contract struct {
	empId    int
	basicPay rune
}

type Employee interface {
	calculateSalary() rune
}

func main() {
	fes := []Freelance{
		{
			empId: 1,
			rate:  7000,
			days:  30,
		},
		{
			empId: 2,
			rate:  5000,
			days:  60,
		},
	}

	pes := []Permanent{
		{
			empId:    3,
			basicPay: 2000,
			pf:       100,
		},
		{
			empId:    4,
			basicPay: 3000,
			pf:       150,
		},
	}

	ces := []Contract{
		{
			empId:    5,
			basicPay: 8000,
		},
		{
			empId:    6,
			basicPay: 10000,
		},
	}

	fmt.Println("Freelance Employees")
	for _, fe := range fes {
		salary := calculateSalary(fe)
		fmt.Printf("Id: %d, salary: Rs. %d \n", fe.empId, salary)
	}

	fmt.Println("\nPermanent Employees")
	for _, pe := range pes {
		salary := calculateSalary(pe)
		fmt.Printf("Id: %d, salary: Rs. %d \n", pe.empId, salary)
	}

	fmt.Println("\nContract Employees")
	for _, ce := range ces {
		salary := calculateSalary(ce)
		fmt.Printf("Id: %d, salary: Rs. %d \n", ce.empId, salary)
	}
}

func calculateSalary(e Employee) rune {
	return e.calculateSalary()
}

func (p Permanent) calculateSalary() rune {
	return p.basicPay + p.pf
}

func (c Contract) calculateSalary() rune {
	return c.basicPay
}

func (f Freelance) calculateSalary() rune {
	return f.rate * f.days
}

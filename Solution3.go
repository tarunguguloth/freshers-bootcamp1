package main

import "fmt"

type Employee interface {
	getSalary() int
}

type fullTimeEmployee struct{
	name string
	no_days int
	salary int
}

func (f fullTimeEmployee) getSalary()int{
	return f.salary * f.no_days
}

type freeLancer struct{
	name string
	no_hrs int
	salary int

}

func (f freeLancer) getSalary() int{
	return f.salary * f.no_hrs
}

type contractor struct{
	name string
	no_days int
	salary int
}

func (c contractor) getSalary() int{
	return c.salary * c.no_days
}

func main(){
	fullEmployee := fullTimeEmployee{"Tarun",20,500}
	contract := contractor{"Varun", 20, 100}
	freelance := freeLancer{"Pavan", 20,10}

	fmt.Println(fullEmployee.getSalary())
	fmt.Println(contract.getSalary())
	fmt.Println(freelance.getSalary())
}

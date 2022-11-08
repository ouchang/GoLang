package main

import "fmt"

type Address struct {
	Street   string
	City     string
	Postcode float64
}

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address //anonimowe pole
}

type Employee struct {
	Name   string
	Salary float64
}

func createAddress(street string, city string, postcode float64) Address {
	return Address{Street: street, City: city, Postcode: postcode}
}

func createSubscriber(name string, street string, city string, postcode float64) Subscriber {
	subscriberAddress := createAddress(street, city, postcode)
	return Subscriber{Name: name, Rate: 15.99, Active: true, Address: subscriberAddress}
}

func (subscriber Subscriber) showInfo() {
	fmt.Printf("Name: %s\n Rate: %.2f\n Active: %t\n ADDRESS\n Street: %s\n City: %s\n Postcode: %f\n\n", subscriber.Name, subscriber.Rate, subscriber.Active, subscriber.Street, subscriber.City, subscriber.Postcode)
}

func applyDiscount(subscriber *Subscriber) {
	subscriber.Rate = 4.99
}

func createEmployee(name string, salary float64) Employee {
	return Employee{Name: name, Salary: salary}
}

func (employee Employee) showInfo() {
	fmt.Printf("Name: %s\n Salary: %f\n\n", employee.Name, employee.Salary)
}

func main() {
	var s1 Subscriber
	var e1 Employee

	s1 = createSubscriber("Nathan Smith", "BlueSky", "Vancouver", 34544)
	s2 := createSubscriber("Susan Wisor", "Marvelous", "Paris", 22438)
	e1 = createEmployee("Michael Scott", 8997)

	s1.showInfo()
	s2.showInfo()
	applyDiscount(&s2)
	s2.showInfo()
	e1.showInfo()
}

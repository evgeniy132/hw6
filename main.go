package main

import "fmt"

type Transport interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

type PassengerTransport interface {
	Transport
	BoardPassengers()
	DisembarkPassengers()
}

type Vehicle struct {
	Name string
}

type Car struct {
	Vehicle
}

func (c Car) Move() {
	fmt.Println("Car is moving")
}

func (c Car) Stop() {
	fmt.Println("Car stopped")
}

func (c Car) ChangeSpeed(speed int) {
	fmt.Printf("Car speed changed to %d\n", speed)
}

func (c Car) BoardPassengers() {
	fmt.Println("Passengers boarded the car")
}

func (c Car) DisembarkPassengers() {
	fmt.Println("Passengers disembarked from the car")
}

type Train struct {
	Vehicle
}

func (t Train) Move() {
	fmt.Println("Train is moving")
}

func (t Train) Stop() {
	fmt.Println("Train stopped")
}

func (t Train) ChangeSpeed(speed int) {
	fmt.Printf("Train speed changed to %d\n", speed)
}

func (t Train) BoardPassengers() {
	fmt.Println("Passengers boarded the train")
}

func (t Train) DisembarkPassengers() {
	fmt.Println("Passengers disembarked from the train")
}

type Plane struct {
	Vehicle
}

func (p Plane) Move() {
	fmt.Println("Plane is flying")
}

func (p Plane) Stop() {
	fmt.Println("Plane landed")
}

func (p Plane) ChangeSpeed(speed int) {
	fmt.Printf("Plane speed changed to %d\n", speed)
}

func (p Plane) BoardPassengers() {
	fmt.Println("Passengers boarded the plane")
}

func (p Plane) DisembarkPassengers() {
	fmt.Println("Passengers disembarked from the plane")
}

type Route struct {
	Transports []Transport
}

func (r *Route) AddTransport(transport Transport) {
	r.Transports = append(r.Transports, transport)
}

func (r Route) ShowTransports() {
	fmt.Println("Transports on the route:")
	for _, transport := range r.Transports {
		fmt.Printf("%T\n", transport)
	}
}

func main() {
	var route Route

	car := Car{Vehicle{Name: "Car"}}
	train := Train{Vehicle{Name: "Train"}}
	plane := Plane{Vehicle{Name: "Plane"}}

	route.AddTransport(car)
	route.AddTransport(train)
	route.AddTransport(plane)

	route.ShowTransports()

	for _, transport := range route.Transports {
		fmt.Printf("Traveling with %T:\n", transport)
		transport.Move()

		if pt, ok := transport.(PassengerTransport); ok {
			pt.BoardPassengers()
			pt.ChangeSpeed(500)
			pt.Stop()
			pt.DisembarkPassengers()
		}

		fmt.Println()
	}
}

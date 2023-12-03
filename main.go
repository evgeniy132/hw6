package main

import "fmt"

type Transport interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

type PassengerTransport interface {
	Transport
	BoardPassengers(int)
	DisembarkPassengers(int)
}

type Vehicle struct {
	Name string
}

func (v Vehicle) Move() {
	fmt.Printf(" %s is moving", v.Name)
}

func (v Vehicle) Stop() {
	fmt.Printf("%s stopped", v.Name)
}

func (v Vehicle) ChangeSpeed(speed int) {
	fmt.Printf(" %s speed changed to %d\n", v.Name, speed)
}

type PassengersVehicle struct {
	PassengerCount int
	Vehicle
}

func (p *PassengersVehicle) BoardPassengers(boardPassengersCount int) {
	p.PassengerCount = p.PassengerCount + boardPassengersCount
	fmt.Printf("Passengers %d boarded the %s, now in %s %d passangers", boardPassengersCount, p.Name, p.Name, p.PassengerCount)
}

func (p *PassengersVehicle) DisembarkPassengers(disembarkPassengers int) {
	if disembarkPassengers > p.PassengerCount {
		diff := disembarkPassengers - p.PassengerCount
		fmt.Printf("Not enought passangers, in %s is %d passangers less than we want\n ", p.Name, diff)
		disembarkPassengers = p.PassengerCount

	}
	p.PassengerCount = p.PassengerCount - disembarkPassengers
	fmt.Printf("Passengers %d disembarked from the %s, now in %s %d passangers\n", disembarkPassengers, p.Name, p.Name, p.PassengerCount)
}

type Car struct {
	*PassengersVehicle
}

const maxPasInCar = 5

func (c Car) BoardPassengers(Bp int) {

	if Bp+c.PassengerCount > maxPasInCar {
		Bp = maxPasInTrain - c.PassengerCount
		fmt.Printf("yoy too many passangers, on board goes only %d passangers", Bp)
	}
	c.PassengersVehicle.BoardPassengers(Bp)
}

type Train struct {
	*PassengersVehicle
}

const maxPasInTrain = 25

func (t Train) BoardPassengers(Bp int) {

	if Bp+t.PassengerCount > maxPasInTrain {
		Bp = maxPasInTrain - t.PassengerCount
		fmt.Printf("yoy too many passangers, on board goes only %d passangers", Bp)
	}
	t.PassengersVehicle.BoardPassengers(Bp)
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
	car := Car{}
	car.PassengersVehicle = &PassengersVehicle{
		Vehicle: Vehicle{Name: "Car"},
	}
	train := Train{}
	train.PassengersVehicle = &PassengersVehicle{
		Vehicle: Vehicle{Name: "Train"},
	}
	plane := Plane{Vehicle{Name: "Plane"}}

	route.AddTransport(car)
	route.AddTransport(train)
	route.AddTransport(plane)

	route.ShowTransports()

	for _, transport := range route.Transports {
		fmt.Printf("Traveling with %T:\n", transport)
		transport.Move()
		transport.ChangeSpeed(100)
		if pt, ok := transport.(PassengerTransport); ok {
			fmt.Println("enter passanges count:")
			var pcount int
			fmt.Scan(&pcount)
			pt.BoardPassengers(pcount)
			pt.Stop()
			fmt.Println("enter go out passanges count :")
			var outcount int
			fmt.Scan(&outcount)
			pt.DisembarkPassengers(outcount)
		}

		fmt.Println()
	}
}

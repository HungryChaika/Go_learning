package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}
type Bicycle struct{}
type Unicycle struct{}

func (c Car) move() {
	fmt.Println("Машина едет")
}

func (b Bicycle) move() {
	fmt.Println("Велик катит")
}

func (u Unicycle) move() {
	fmt.Println("Как этот моноцикл вообще едет? Но всё же... Едет!")
}

func main() {
	var (
		car      Vehicle = Car{}
		bicycle  Vehicle = Bicycle{}
		unicycle Vehicle = Unicycle{}
		car1     Car     = Car{}
	)
	car.move()
	bicycle.move()
	unicycle.move()
	fmt.Println("=====================\n=====================")
	drive(car1)
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

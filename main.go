package main

import "fmt"

type Car struct {
	PlateNumber string
	Color       string
}

func newCar(plateNumber string, color string) *Car {
	return &Car{PlateNumber: plateNumber, Color: color}
}

type Parking map[int]*Car

func AllocatePark(slot int) Parking {
	parking := make(Parking)
	for i := range slot {
		parking[i] = nil
	}
	return parking
}

func (p Parking) status() {
	fmt.Println("Slot Registration Parking:")
	fmt.Println("No\t\tPlate Number\t\tColor")
	for i := 1; i <= len(p); i++ {
		car := p[i-1]
		if car != nil {
			fmt.Printf("%d.\t\t%s\t\t%s\n", i, car.PlateNumber, car.Color)
		} else {
			fmt.Printf("%d.\t\tFree Slot\n", i)
		}
	}
}

func (p Parking) parkCar(plateNumber string, color string) {
	for i := 1; i <= len(p); i++ {
		car := p[i-1]
		if car != nil && car.PlateNumber == plateNumber {
			fmt.Printf("Car %s is already parked at slot %d\n", plateNumber, i)
			return
		}
		if car == nil {
			p[i-1] = newCar(plateNumber, color)
			fmt.Printf("slot allocated at slot: %d\n", i)
			return
		}
	}
	fmt.Println("parking not available")
}

func (p Parking) unparkCar(slot, hours int) error {
	if slot < 1 || slot > len(p) {
		return fmt.Errorf("slot not found")
	}
	car := p[slot-1]
	if car == nil {
		return fmt.Errorf("slot already empty")
	}

	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}

	fmt.Printf("Registration number %s with slot number %d is free with Charge $%d\n",
		car.PlateNumber, slot, charge)
	p[slot-1] = nil
	return nil
}

func main() {
	var capacity int
	fmt.Println("capacity:")
	fmt.Scan(&capacity)

	parking := AllocatePark(capacity)

	for {
		fmt.Println("1 park - Input 'PLATE NUMBER(space)COLOR'")
		fmt.Println("2 leave - Input 'SLOT NUMBER(space)HOURS PARKED'")
		fmt.Println("3 status")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			var plate, color string
			fmt.Scan(&plate, &color)
			parking.parkCar(plate, color)
		}

		if choice == 2 {
			var slot, hours int
			fmt.Scan(&slot, &hours)
			parking.unparkCar(slot, hours)
		}

		if choice == 3 {
			parking.status()
		}
	}
}

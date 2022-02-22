package main

import (
	"fmt"

	"github.com/Golang/cars/machine"
	"github.com/Golang/cars/client"

)

func main() {
	cfgBmw := machine.Config{
		Power:              150,
		DoorsCount:         5,
		MarkCar:            "BMW",
		ModelCar:           "5 series",
		YearManufactureCar: 1997,
		Price:              3000000,
		StatusBuyCar:       true,
	}

	bmw := machine.NewcInfoCar(cfgBmw)
	fmt.Println(bmw)


	client.NewClient(client.Config{
		Car: bmw,
	})


}

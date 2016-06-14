package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// SwampCooler interface
type SwampCooler interface {
	Close()
	Open() error
	ResetPins()
	SetPump(bool)
}

// GPIOCooler implementation
type gpioCooler struct {
}

func (c gpioCooler) Close() {
	fmt.Println("Closing...")
}

func (c gpioCooler) Open() error {
	fmt.Println("Opening...")
	return rpio.Open()
}

func (c gpioCooler) ResetPins() {
	fmt.Println("Reseting Pins")
}

func (c gpioCooler) SetPump(b bool) {

}

type mockCooler struct {
	motor     bool
	motorHigh bool
	pump      bool
}

func (c mockCooler) Close() {
	fmt.Println("Closing...")
}

func (c mockCooler) Open() error {
	fmt.Println("Opening...")
	return nil
}

func (c mockCooler) ResetPins() {
	fmt.Println("Reseting Pins")
}

func (c mockCooler) SetPump(b bool) {
	c.motor = b
	c.pump = b
	c.printState()
}

func (c mockCooler) printState() {
	fmt.Printf("motor: %t, motorHigh: %t, pump: %t\n", c.motor, c.motorHigh, c.pump)
}

const (
	pump      = rpio.Pin(4) // relay 1
	motor     = rpio.Pin(3) // relay 2
	motorHigh = rpio.Pin(2) // relay 3
)

func main() {
	var cooler SwampCooler = mockCooler{}

	if err := cooler.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cooler.Close()

	// Open and map memory to access gpio, check for errors

	//
	// defer rpio.Close()

	// pump.Output()
	// motor.Output()
	// motorHigh.Output()
	//
	for x := 0; x < 3; x++ {
		cooler.SetPump(true)
		time.Sleep(time.Second)
		cooler.SetPump(false)
		time.Sleep(time.Second * 3)
	}

	cooler.ResetPins()
	//
	// pump.High()
	// motor.High()
	// motorHigh.High()
}

package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// SwampCooler interface
type SwampCooler interface {
	Close()
	GetMotor() bool
	GetPump() bool
	Open() error
	ResetPins()
	SetMotor(bool)
	SetPump(bool)
}

const (
	pump  = rpio.Pin(4) // relay 1
	motor = rpio.Pin(3) // relay 2
	// motorHigh = rpio.Pin(2) // relay 3
)

func setRelay(pin rpio.Pin, b bool) {
	if b {
		pin.Low()
	} else {
		pin.High()
	}
}

type logCooler struct {
	motor bool
	pump  bool
}

func (c logCooler) Close() {
	log.Println("Closing...")
}

func (c logCooler) GetMotor() bool {
	return c.motor
}

func (c logCooler) GetPump() bool {
	return c.pump
}

func (c logCooler) Open() error {
	log.Println("Opening...")
	return nil
}

func (c logCooler) ResetPins() {
	log.Println("Reseting Pins...")
}

func (c logCooler) SetMotor(b bool) {
	c.motor = b
	c.logState()
}

func (c logCooler) SetPump(b bool) {
	c.pump = b
	c.logState()
}

func (c logCooler) logState() {
	log.Printf("motor: %t, pump: %t\n", c.GetMotor(), c.GetPump())
}

type gpioCooler struct {
	log logCooler
}

func (c gpioCooler) Close() {
	c.log.Close()

	rpio.Close()
}

func (c gpioCooler) GetMotor() bool {
	return c.log.motor
}

func (c gpioCooler) GetPump() bool {
	return c.log.pump
}

func (c gpioCooler) Open() error {
	c.log.Open()

	if err := rpio.Open(); err != nil {
		return err
	}

	pump.Output()
	motor.Output()

	return nil
}

func (c gpioCooler) ResetPins() {
	c.log.ResetPins()

	c.SetMotor(false)
	c.SetPump(false)
}

func (c gpioCooler) SetMotor(b bool) {
	c.log.SetMotor(b)

	setRelay(motor, b)
}

func (c gpioCooler) SetPump(b bool) {
	c.log.SetPump(b)

	setRelay(pump, b)
}

func main() {
	var cooler SwampCooler = gpioCooler{}
	// var cooler SwampCooler = logCooler{}

	if err := cooler.Open(); err != nil {
		log.Fatal(err)
	}
	defer cooler.Close()

	for x := 0; x < 3; x++ {
		cooler.SetPump(true)
		time.Sleep(time.Second)
		cooler.SetPump(false)
		time.Sleep(time.Second * 3)
	}

	cooler.ResetPins()
}

package main

import (
	"log"
	"time"
)

func main() {
	var cooler SwampCooler = gpioCooler{}
	// var cooler SwampCooler = logCooler{}

	if err := cooler.Open(); err != nil {
		log.Fatal(err)
	}
	defer cooler.Close()

	cooler.ResetPins()

	for x := 0; x < 3; x++ {
		cooler.SetPump(true)
		time.Sleep(time.Second)
		cooler.SetPump(false)
		time.Sleep(time.Second * 3)
	}

	cooler.ResetPins()
}

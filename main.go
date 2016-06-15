package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

// FIXME insane code :P

func main() {
	var cooler SwampCooler = &gpioCooler{}
	// var cooler SwampCooler = &logCooler{}
	var coolerMutex = &sync.Mutex{}

	if err := cooler.Open(); err != nil {
		log.Fatal(err)
	}
	defer cooler.Close()

	cooler.ResetPins()

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		coolerMutex.Lock()
		defer coolerMutex.Unlock()

		if cooler.GetPump() {
			cooler.SetPump(false)
			cooler.SetMotor(false)
		} else {
			cooler.SetPump(true)
			time.Sleep(1 * time.Second)
			cooler.SetMotor(true)
		}

		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		coolerMutex.Lock()
		defer coolerMutex.Unlock()

		t, err := template.New("home.html").Parse(`
			<!DOCTYPE html>
			<html>
			  <head>
			    <meta charset="utf-8">
					<meta name="viewport" content="width=device-width, initial-scale=1">
			    <title>SwampCooler</title>
			  </head>
			  <body>
					<form action="/toggle" method="post" style="max-width: 300px; height: 100px; margin: 20px auto;">
					  <input style="border: none; height: 100px; width: 100%;" type="submit" name="name" value="{{.}}">
					</form>
			  </body>
			</html>
		`)
		if err != nil {
			log.Fatal(err)
		}

		var buttonText string
		if cooler.GetPump() {
			buttonText = "Turn Off"
		} else {
			buttonText = "Turn On"
		}

		t.Execute(w, buttonText)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

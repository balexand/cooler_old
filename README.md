# cooler

Raspberry Pi swamp cooler controller.

## Hardware

Connect [like this](https://www.youtube.com/watch?v=OQyntQLazMU). GPIO pin number constants are near the top of main.go.

## Build

```
GOOS=linux GOARCH=arm go build -v
```

Then run the resulting binary on the Pi.

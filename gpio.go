package main

import (
	"gobot.io/x/gobot/platforms/raspi"
)

func getGpio(pinNumber string) (value byte) {
	value, err := raspi.NewAdaptor().DigitalRead(pinNumber)
	fatal(err)
	return value
}

func setGpio(pinNumber string, value byte)  {
	fatal(raspi.NewAdaptor().DigitalWrite(pinNumber, value))
}

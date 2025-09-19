package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/filosocode/enerbit_golang/device"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	data := device.SimulateConsumption("device-123")
	fmt.Printf("Consumo: %+v\n", data)
}

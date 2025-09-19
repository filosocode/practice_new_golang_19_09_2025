package main

import (
	"math/rand"
	"time"

	"github.com/filosocode/enerbit_golang/producer"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := producer.NewProducer("127.0.0.1:6381", "consumption_channel", "device-001")
	p.Start(2 * time.Second)
}

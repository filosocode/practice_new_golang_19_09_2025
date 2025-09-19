package producer

import (
	"context"
	"encoding/json"
	"log"
	"math/rand/v2"
	"time"

	"github.com/filosocode/enerbit_golang/device"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Producer struct {
	Client   *redis.Client
	Channel  string
	DeviceID string
}

func NewProducer(redisAddr, channel, deviceID string) *Producer {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &Producer{
		Client:   rdb,
		Channel:  channel,
		DeviceID: deviceID,
	}
}

func (p *Producer) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for t := range ticker.C {
		data := device.ConsumptionData{
			DeviceID:  p.DeviceID,
			Timestamp: t,
			Value:     0.5 + rand.Float64()*1.5,
		}

		payload, err := json.Marshal(data)
		if err != nil {
			log.Println("error serializando JSON:", err)
			continue
		}

		err = p.Client.Publish(ctx, p.Channel, payload).Err()
		if err != nil {
			log.Println("error publicando en Redis:", err)
		}
	}
}

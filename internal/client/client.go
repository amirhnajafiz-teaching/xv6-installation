package client

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/official-stallion/race/internal/metrics"

	sdk "github.com/official-stallion/go-sdk"
)

type Client struct {
	Metrics *metrics.Handler
	Debug   bool
}

func (c Client) Consumer(host, topic string) {
	var wg sync.WaitGroup

	client, err := sdk.NewClient(host)
	if err != nil {
		log.Println(fmt.Errorf("failed to connect error=%w", err))

		return
	}

	wg.Add(1)
	c.Metrics.Subscribe(topic)

	client.Subscribe(topic, func(bytes []byte) {
		c.Metrics.Receive(topic)
		if c.Debug {
			log.Println(fmt.Sprintf("got %d bytes", len(bytes)))
		}
	})

	wg.Wait()
}

func (c Client) Provider(host, topic, message string, wait int) {
	client, err := sdk.NewClient(host)
	if err != nil {
		log.Println(fmt.Errorf("failed to connect error=%w", err))

		return
	}

	for {
		if er := client.Publish(topic, []byte(message)); er != nil {
			c.Metrics.Failed()
			if c.Debug {
				log.Println(fmt.Errorf("failed to publish erorr=%w", er))
			}
		} else {
			c.Metrics.Publish(topic)
		}

		time.Sleep(time.Duration(wait) * time.Second)
	}
}

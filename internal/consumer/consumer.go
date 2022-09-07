package consumer

import (
	"fmt"

	"github.com/amirhnajafiz/stallion"
)

type Consumer struct {
	Topic string
}

func (c Consumer) Start(addr string) error {
	conn, err := stallion.NewClient(addr)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	conn.Subscribe(c.Topic, func(bytes []byte) {
		fmt.Println(string(bytes))
	})

	select {}
}

package provider

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/amirhnajafiz/stallion"
)

type Provider struct {
	Data   []byte
	Number int
	Timout int
	Topics []string
}

func (p Provider) Start(addr string) error {
	c, err := stallion.NewClient(addr)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	for i := 0; i < p.Number; i++ {
		index := rand.Int() % len(p.Topics)

		err = c.Publish(p.Topics[index], p.Data)
		if err != nil {
			return fmt.Errorf("publish failed: %w", err)
		}

		time.Sleep(time.Duration(p.Timout) * time.Millisecond)
	}

	return nil
}

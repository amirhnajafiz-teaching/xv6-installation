package client

import (
	"fmt"
	"log"
	"sync"

	sdk "github.com/official-stallion/go-sdk"
)

func Consumer(host, topic string) {
	var wg sync.WaitGroup

	client, err := sdk.NewClient(host)
	if err != nil {
		log.Println(fmt.Errorf("failed to connect error=%w", err))
	}

	wg.Add(1)

	client.Subscribe(topic, func(bytes []byte) {
		log.Println(fmt.Sprintf("got %d bytes", len(bytes)))
	})

	wg.Wait()
}

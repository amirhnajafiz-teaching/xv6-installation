package provider

import "github.com/amirhnajafiz/stallion"

var (
	topic = "topic"
	data  = []byte("data")
)

func Start(addr string) {
	c, err := stallion.NewClient(addr)
	if err != nil {
		panic(err)
	}

	err = c.Publish(topic, data)
	if err != nil {
		panic(err)
	}
}

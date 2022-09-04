package main

import "github.com/amirhnajafiz/stallion"

func main() {
	if err := stallion.NewServer(":9254"); err != nil {
		panic(err)
	}
}

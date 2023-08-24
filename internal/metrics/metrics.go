package metrics

import "sync"

type (
	Handler struct {
		Metrics Metrics
		lock    sync.Mutex
	}

	Metrics struct {
		Providers       int            `json:"providers"`
		Consumers       int            `json:"consumers"`
		FailedPublish   int            `json:"failed_publish"`
		Publish         map[string]int `json:"publish"`
		Subscribe       map[string]int `json:"subscribe"`
		ReceivedMessage map[string]int `json:"received_message"`
		Topics          []string       `json:"topics"`
	}
)

func (h *Handler) Publish(topic string) {
	h.lock.Lock()
	h.Metrics.Publish[topic]++
	h.lock.Unlock()
}

func (h *Handler) Failed() {
	h.lock.Lock()
	h.Metrics.FailedPublish++
	h.lock.Unlock()
}

func (h *Handler) Subscribe(topic string) {
	h.lock.Lock()
	h.Metrics.Subscribe[topic]++
	h.lock.Unlock()
}

func (h *Handler) Receive(topic string) {
	h.lock.Lock()
	h.Metrics.ReceivedMessage[topic]++
	h.lock.Unlock()
}

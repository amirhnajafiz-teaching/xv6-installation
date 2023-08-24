package metrics

type Metrics struct {
	Providers       int            `json:"providers"`
	Consumers       int            `json:"consumers"`
	FailedPublish   int            `json:"failed_publish"`
	Publish         map[string]int `json:"publish"`
	Subscribe       map[string]int `json:"subscribe"`
	ReceivedMessage map[string]int `json:"received_message"`
	Topics          []string       `json:"topics"`
}

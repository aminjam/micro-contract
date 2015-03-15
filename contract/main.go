package contract

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bitly/go-nsq"
)

type Message struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

type Model struct {
	Pattern string `json:"pattern"`
	Call    func([]byte)
}

var Collection map[string]Model

func Handler() nsq.HandlerFunc {
	return nsq.HandlerFunc(func(m *nsq.Message) error {
		var decoded Message
		fmt.Println("Contract:", string(m.Body[:]))
		err := json.Unmarshal(m.Body, &decoded)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing the json: %s\n", err.Error())
			return nil
		}
		contract, ok := Collection[decoded.Key]
		if ok {
			contract.Call(decoded.Value)
		}
		return nil
	})
}

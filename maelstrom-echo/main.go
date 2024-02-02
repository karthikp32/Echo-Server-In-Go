package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {

	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		// Declare a map with key-value pairs of type {string:any}
		var requestBody map[string]any

		//"unmarshal" or convert the message body's json AKA deserializing
		//and return the error data structure stored in variable err
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}

		requestBody["type"] = "echo_ok"

		return n.Reply(msg, requestBody)

		

	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

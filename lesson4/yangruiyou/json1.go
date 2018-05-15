package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//

	doc := []byte(`{"id":12345,"name":"Test doc","payload":{\"message\":\"test\"}}`)
	var entry LogEntry
	if err := json.Unmarshal(doc, &entry), err != nil {
		fmt.Println("error", err)

	}
	fmt.Printf("%v", entry)
}

type LogEntry struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type LogPayload struct {
	Message string `json:"message"`
}

type fauxLogPayload LogPayload

func (lp *LogPayload) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err

	}
	var f fauxLogPayload
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}

	*lp = LogPayload(f)

	return nil

}

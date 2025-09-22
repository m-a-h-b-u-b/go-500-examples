package main

import (
	"encoding/json"
	"fmt"
)

type Example struct {
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"name":"Go"}`)
	var ex Example
	json.Unmarshal(data, &ex)
	fmt.Println(ex.Name)
}

package main

import (
	"encoding/json"
	"fmt"
)

// Custom struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// Example struct instance
	person := Person{
		Name: "John",
		Age:  30,
		City: "New York",
	}

	// Marshal the struct into a map
	var data map[string]interface{}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if a key exists
	keyToCheck := "age"
	if value, ok := data[keyToCheck]; ok {
		fmt.Printf("Key '%s' exists, and its value is: %v\n", keyToCheck, value)
	} else {
		fmt.Printf("Key '%s' does not exist\n", keyToCheck)
	}
}

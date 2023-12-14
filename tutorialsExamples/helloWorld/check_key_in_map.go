package main

import "fmt"

func main() {
	// check --> https://golang.cafe/blog/how-to-check-if-a-map-contains-a-key-in-go
	var exist bool
	var evenType string
	a := map[string]string{
		"eventType": "Dani",
	}
	evenType, exist = a["eventType"]
	if exist {
		fmt.Println("Existe!")
		fmt.Println(evenType)
	} else {
		fmt.Println("No existe!")
		//fmt.Println(evenType)
	}
}

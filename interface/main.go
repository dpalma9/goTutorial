package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getPrometheusAlerts(url string, token string) {
	client := &http.Client{}

	//Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	var bearer = "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("IBMInstanceID", "myID")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(string(data))

	// Body response to interface
	// On this proposal, instead of using struct, we use interface. And, in this case, our interface variable is a map of strings interface:
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct
		log.Fatal("Can not unmarshal JSON")
	}

	//log.Println(result)
	// In order to print interface fields, it has to be trated like a map var.
	// check: https://go.dev/blog/json
	if status, err := result["status"].(string); err {
		log.Println("Status --> ", status)
	}
}

func main() {
	getPrometheusAlerts("https://private.eu-de.monitoring.cloud.ibm.com/prometheus/api/v1/alerts", "mytokenlargocifradosuperseguro")
}

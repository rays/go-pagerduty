package main

import (
	"github.com/rays/go-pagerduty"
	"log"
)

func main() {
	var key = "PagerDuty API Key"
	event := pagerduty.Event{
		Type:        "trigger",
		ServiceKey:  key,
		Description: "Example event",
	}
	resp, err := pagerduty.CreateEvent(event)
	if err != nil {
		log.Println(resp)
		log.Fatalln("ERROR:", err)
	}
	log.Println("Incident key:", resp.IncidentKey)
}

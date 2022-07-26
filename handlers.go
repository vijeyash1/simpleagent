package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (app *jsPool) GitHandler(c *gin.Context) {
	// Get the message from the request body

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}
	metricsJson, _ := json.Marshal(payload)
	// Publish the message to the stream
	if _, err := app.js.Publish(eventSubject, metricsJson); err != nil {
		log.Println("Error publishing to stream:", err)
		return
	}
	fmt.Println(string(metricsJson))
	log.Printf("Metrics with eventSubject:%s has been published\n", eventSubject)

}

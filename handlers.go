package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func (app *jsPool) GitHandler(c *gin.Context) {
	// Get the message from the request body

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}

	// Publish the message to the stream
	if _, err := app.js.Publish(eventSubject, payload); err != nil {
		log.Println("Error publishing to stream:", err)
		return
	}
	fmt.Println(string(payload))
	log.Printf("Metrics with eventSubject:%s has been published\n", eventSubject)

}

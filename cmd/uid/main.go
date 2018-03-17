package main

import (
	"github.com/codehakase/uid"
	"log"
)

func main() {
	randID := uid.New(30)

	log.Printf("UID generated: %+v", randID)
}

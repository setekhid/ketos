package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	err := Command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

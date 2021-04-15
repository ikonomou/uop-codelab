package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Started....")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			log.Errorf("I=%d", i)
		} else {
			log.Warnf("I=%d", i)
		}
	}
}

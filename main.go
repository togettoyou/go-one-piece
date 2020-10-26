package main

import (
	"go-one-piece/util"
	"go-one-piece/util/conf"
	"go-one-piece/util/logging"
)

func init() {
	conf.Setup()
	logging.Setup()
}

func main() {
	reload := make(chan int, 1)
	conf.OnConfigChange(func() { reload <- 1 })
	for {
		select {
		case <-reload:
			util.Reset()
			logging.Get().Infoln("OnConfigChange")
		}
	}
}

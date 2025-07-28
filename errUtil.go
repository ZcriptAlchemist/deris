package main

import "log"

func PanicOnError(err error, msg string) {
	log.Println(err)
	log.Panicln(msg)
}

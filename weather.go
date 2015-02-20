package main

import (
	"fmt"
	"log"

	"github.com/tarm/goserial"
)

func main() {
	fmt.Println("hello")
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		return
	}
	_ = s
}

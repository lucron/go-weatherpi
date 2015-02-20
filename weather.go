package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/tarm/goserial"
)

func main() {

	u := CreateOrOpenDB("weather.rrd")
	_ = u
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(s)
	for {
		data, err := reader.ReadString('\n')
		data = strings.Trim(data, "\r\n")
		if err != nil {
			log.Println(err)
		}
		log.Println("got data:", data)
	}

}

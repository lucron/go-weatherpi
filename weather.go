package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/tarm/goserial"
	"github.com/ziutek/rrd"
)

func main() {
	var wg sync.WaitGroup
	u := CreateOrOpenDB("weather.rrd")
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bufio.NewReader(s)
	wg.Add(1)
	go ReadAndWriteData(u, reader)
	_ = exportData("weather.rrd")
	http.HandleFunc("/", serve)
	http.ListenAndServe(":80", nil)
	wg.Wait()
}

func ReadAndWriteData(u *rrd.Updater, reader *bufio.Reader) {
	for {
		data, err := reader.ReadString('\n')
		data = strings.Trim(data, "\r\n")
		if err != nil {
			log.Println(err)
		}
		log.Println("got data:", data)
		writeData(u, data)
	}
}

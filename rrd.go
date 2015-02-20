package main

import (
	"log"
	"os"
	"time"

	"github.com/ziutek/rrd"
)

// returns an *rrd.Updater. Creates a new empty rrd if necessary.
func CreateOrOpenDB(f string) *rrd.Updater {
	if _, err := os.Stat(f); err == nil {
		log.Println("pened existing db", f)
		return rrd.NewUpdater(f)
	}
	c := rrd.NewCreator(f, time.Now(), 300)
	c.DS("temp1", "GAUGE", "600", -40, 50)
	c.DS("hum1", "GAUGE", "600", 0, 100)
	//1 day - 5min resolution
	c.RRA("AVERAGE", 0.5, 1, 288)
	c.RRA("MAX", 0.5, 1, 288)
	c.RRA("MIN", 0.5, 1, 288)
	//1 week - 15min resolution
	c.RRA("AVERAGE", 0.5, 3, 672)
	c.RRA("MAX", 0.5, 3, 672)
	c.RRA("MIN", 0.5, 3, 672)
	//1 year - 1h resolution
	c.RRA("AVERAGE", 0.5, 12, 8880)
	c.RRA("MAX", 0.5, 12, 8880)
	c.RRA("MIN", 0.5, 12, 8880)
	err := c.Create(false)
	if err != nil {
		log.Println("Error creating RRD:", err.Error())
	}
	log.Println("created new db", f)
	return rrd.NewUpdater(f)
}

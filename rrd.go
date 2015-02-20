package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ziutek/rrd"
)

// returns an *rrd.Updater. Creates a new empty rrd if necessary.
func CreateOrOpenDB(f string) *rrd.Updater {
	if _, err := os.Stat(f); err == nil {
		log.Println("opened existing db", f)
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

func writeData(u *rrd.Updater, raw string) {
	values := strings.Split(raw, ";")
	//values are 3*n for temp, 11+n for hum
	log.Println("temp1: %s - hum1: %s\n", values[3], values[11])
	//rrd doesnt like ","
	err := u.Update(time.Now(), strings.Replace(values[3], ",", ".", -1), values[11])
	if err != nil {
		log.Println("Error updating DB:", err.Error())
	}
}

func exportData(f string) []byte {
	out, err := exec.Command("rrdtool", "xport", "-s", "now-48h", "--step", "300", "DEF:a="+f+":temp1:AVERAGE", "XPORT:a:\"moep\"").CombinedOutput()
	if err != nil {
		log.Println("Error exporting data:", err.Error())
	}
	return out
}

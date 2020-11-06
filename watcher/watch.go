package watcher

import (
	"fmt"
	"time"

	"github.com/darkLord19/dfon/db"
	"github.com/darkLord19/dfon/parser"
)

func startWatcher(dbase *parser.Database) {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				cols, err := db.GetColumsAboveThreshold(dbase)
				if err != nil {
					panic("something bad happened")
				}
				fmt.Println(cols)
			}
		}
	}()
}

// Watch starts go routines to monitor given databases
func Watch(c *parser.Config) {
	for _, d := range c.DBList {
		startWatcher(&d)
	}
}

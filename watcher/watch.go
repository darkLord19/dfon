package watcher

import (
	"fmt"
	"time"

	"github.com/darkLord19/dfon/parser"
)

func startPostgresWatcher(dbase parser.Database) {
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Println("Hello world")
			}
		}
	}()
}

// Watch starts go routines to monitor given databases
func Watch(c parser.Config) {
	for _, d := range c.DBList {
		switch d.Type {
		case "postgres":
			startPostgresWatcher(d)
		case "mysql":
			startMySQLWatcher(d)
		}
	}
}

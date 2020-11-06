package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/darkLord19/dfon/parser"
	"github.com/darkLord19/dfon/watcher"
)

const (
	defaultFileName = "dfon.json"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// check if default config file exist
func defaultConfigFileExist() bool {
	return fileExists(defaultFileName)
}

func main() {
	configFile := flag.String("c", "", "path to config file")
	flag.Parse()
	if *configFile == "" && !defaultConfigFileExist() {
		fmt.Println("provide path to config file or create file named \"dfon\" with your config in it")
		os.Exit(1)
	} else if !fileExists(*configFile) {
		fmt.Println(*configFile, ": file doesn't exist")
	}
	conf, err := parser.LoadConfig(*configFile)
	if err != nil {
		fmt.Println("error loading config")
	}
	fmt.Println(*configFile, conf)
	watcher.Watch(&conf)
	wait := make(chan struct{})
	<-wait
}

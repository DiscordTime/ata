package main

import (
	"flag"
	"fmt"
	"os"
)

var Version = "0.0.0"

func version() {
	fmt.Printf("Go Test Log Version: %s\n", Version)
}

func main() {
	log_flag := flag.Bool("log", false, "Enable log analisis")
	adb_flag := flag.Bool("adb", false, "Enable adb analisis")
	show_version := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *show_version {
		version()
		os.Exit(0)
	}

	if *log_flag {
		fmt.Println("Run log functions")
	}
	if *adb_flag {
		fmt.Println("Run adb functions")
	}
}

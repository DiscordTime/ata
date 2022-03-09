package main

import (
	"fmt"
	"os"
)

var Version = "0.0.0"

func help() {
	fmt.Println("Usage:")
	fmt.Printf("    %s [ help | version | adb | log ]\n", os.Args[0])
	fmt.Println("\t -h | help\tShow help message")
	fmt.Println("\t -v | version\tShow current version")
	fmt.Println("\t adb\t\tRun adb analisis")
	fmt.Println("\t log\t\tRun log analisis")
	fmt.Println()
	fmt.Println("    Eg:")
	fmt.Printf("\t%s help\n", os.Args[0])
	fmt.Printf("\t%s -h\n", os.Args[0])
	fmt.Printf("\t%s version\n", os.Args[0])
	fmt.Printf("\t%s -v\n", os.Args[0])
	fmt.Printf("\t%s log\n", os.Args[0])
	fmt.Printf("\t%s adb\n", os.Args[0])
}

func version() {
	fmt.Printf("Go Test Log Version: %s\n", Version)
}

func invalidParam(param string) {
	fmt.Printf("Invalid Param: %s\n", param)
	fmt.Printf("Try: %s help\n", os.Args[0])
}

func main() {
	log_flag := false
	adb_flag := false
	argv := os.Args[1:]
	for _, item := range argv {
		switch item {
		case "version", "-v":
			version()
			os.Exit(0)
		case "help", "-h":
			help()
			os.Exit(0)
		case "log":
			log_flag = true
		case "adb":
			adb_flag = true
		default:
			invalidParam(item)
			os.Exit(1)
		}
	}
	if log_flag {
		fmt.Println("Run log functions")
	}
	if adb_flag {
		fmt.Println("Run adb functions")
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func help() {
	fmt.Println("Usage:")
	fmt.Printf("    %s [ help | version ]\n", os.Args[0])
	fmt.Println("\t -h | help\tShow help message")
	fmt.Println("\t -v | version\tShow current version")
	fmt.Println()
	fmt.Println("    Eg:")
	fmt.Printf("\t%s help\n", os.Args[0])
	fmt.Printf("\t%s -h\n", os.Args[0])
	fmt.Printf("\t%s version\n", os.Args[0])
	fmt.Printf("\t%s -v\n", os.Args[0])
}

func version() {
	version := os.Getenv("GO_TEST_LOG_VERSION")
	if version == "" {
		version = "0.0.1"
	}
	fmt.Printf("Go Test Log Version: %s\n", version)
}

func loadEnv() {
	godotenv.Load()
	if os.Getenv("GO_TEST_LOG_VERSION") == "" {
		fmt.Println("App version not found")
		fmt.Println("Try set GO_TEST_LOG_VERSION on .env file, or set this variable on environment")
		os.Exit(1)
	}
}

func invalidParam(param string) {
	fmt.Printf("Invalid Param: %s\n", param)
	fmt.Printf("Try: %s help\n", os.Args[0])
}

func main() {
	loadEnv()
	argv := os.Args[1:]
	for _, item := range argv {
		switch item {
		case "version", "-v":
			version()
			os.Exit(0)
		case "help", "-h":
			help()
			os.Exit(0)
		default:
			invalidParam(item)
		}
	}
}

package main

import (
	"log"
	"os"
)

func main() {
	// Get command line arguments
	args := os.Args
	if len(args) == 1 {
		println("Please supply a file name")
		os.Exit(1)
	}
	fileName := args[1]
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	filePath := currentDir + "/" + fileName
	// Check if the file exists
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		println("File already exists")
		os.Exit(1)
	}
	// Create the file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

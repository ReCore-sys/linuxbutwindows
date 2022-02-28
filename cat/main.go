package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

func notation(size int64) string {
	// Get the largest notation for a byte size in kb/mb/gb/ect.
	var notation string
	var bcount float64
	if size < 1024 {
		notation = "bytes"
		bcount = float64(size)
	} else if size < 1024*1024 {
		notation = "KB"
		bcount = float64(size) / 1024
	} else if size < 1024*1024*1024 {
		notation = "MB"
		bcount = float64(size) / 1024 / 1024
	} else if size < 1024*1024*1024*1024 {
		notation = "GB"
		bcount = float64(size) / 1024 / 1024 / 1024
	} else {
		notation = "TB"
		bcount = float64(size) / 1024 / 1024 / 1024 / 1024
	}
	return fmt.Sprintf("%.2f %s", bcount, notation)
}

func main() {
	// When the file is run, the user will give us a path to a file.
	// We will read the file and print it to the screen.
	// We will also print the file's size to the screen.
	// Open the file.
	if len(os.Args) == 1 {
		fmt.Println("Please provide a file path.")
		os.Exit(1)
	} else if len(os.Args) > 1 {
		path := os.Args[1]
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		stats, err := file.Stat()
		if err != nil {
			log.Println(err)
		}
		// Close the file when we're done.
		defer file.Close()

		// Read the file.
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		print(red + "--- File Contents ---\n" + reset)
		// Print the file's contents as a string to the screen.
		fmt.Println(string(bytes))
		print(red + "--- File Data ---\n" + reset)
		// Size
		size := stats.Size()
		notation := notation(size)

		fmt.Println("Size:", notation)
		// Last Modified
		print("Last Modified: ")
		fmt.Println(stats.ModTime().Format("Mon Jan _2 15:04:05 2006"))
		// File Permissions
		println("Permissions: ", stats.Mode().String())

		print(red + "--- End of File ---\n" + reset)
	}
}

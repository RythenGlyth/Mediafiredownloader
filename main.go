package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var mediafolders []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":

			fmt.Println("Usage: ", os.Args[0], " [OPTIONS]... [MEDIAFIREFOLDER]...")
			fmt.Println()
			fmt.Println("standard input is read as a batch file (see below)")
			fmt.Println()
			fmt.Println("MEDIAFIREFOLDER can be a folder or a file link")
			fmt.Println()

			fmt.Println("Options:")
			fmt.Println("  -h, --help:                  Show this help message")
			fmt.Println("  -b, --batch [FILE]:          Download all folder and file links listed in FILE")
			fmt.Println("  -o, --output [PATH]:         Set the output folder")

			os.Exit(0)
		case "-b", "--batch":
			if i+1 < len(args) {
				i++
				file, err := os.Open(args[i])
				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					mediafolders = append(mediafolders, scanner.Text())
				}
			}
		default:
			mediafolders = append(mediafolders, args[i])
		}
	}

	if isInputFromPipe() {
		f := os.Stdin
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			mediafolders = append(mediafolders, scanner.Text())
		}
	}
	if len(mediafolders) < 1 {
		fmt.Println("You need to specify at least one mediafire folder or file")
		os.Exit(1)
	}
	fmt.Println("Downloading: ", mediafolders)
}
func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

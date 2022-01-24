package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var realArgs []string
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":

			fmt.Println("Mediafire-Downloader")
			fmt.Println()

			fmt.Println("Options:")
			fmt.Println("  -h, --help:          Show this help message")
			fmt.Println()

			fmt.Println("Arguments:")
			fmt.Println("  FOLDERS:             Mediafire folder ID")

			os.Exit(0)
		case "-b", "--batch":

		default:
			realArgs = append(realArgs, arg)
		}
	}

	var f *os.File
	var err error

	if isInputFromPipe() {
		f = os.Stdin
	} else if len(realArgs) < 1 {
		fmt.Println("Help: ", os.Args[0], "--help")
		os.Exit(1)
	} else {
		f, err = os.Open(realArgs[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	sc := bufio.NewScanner(f)

	sc.Scan()

}
func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

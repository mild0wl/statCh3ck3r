package main

import (
	"fmt"
	"os"
	"statCh3ck3r/cmd"
	"strings"
	"sync"
)

func main() {
	fmt.Println("statCh3ck3r.....")

	// input

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	filename := os.Args[1]

	fileData, err := cmd.FileInput_func(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}

	for _, url := range strings.Fields(string(fileData)) {
		wg.Add(1)
		go cmd.StatCheck(url, wg)
	}
	wg.Wait()

}

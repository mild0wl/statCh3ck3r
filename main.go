package main

import (
	"statCh3ck3r/cmd"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("statCh3ck3r.....")
	fileData, err := cmd.FileInput_func("file.txt")

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

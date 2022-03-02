package cmd

import (
	"fmt"
	"net/http"
	"sync"
)

func StatCheck(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	// c <- resp.StatusCode
	fmt.Println(resp.StatusCode, url)
	// close(c)

}

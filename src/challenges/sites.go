package challenge

import (
	"fmt"
	"time"
)

func siteSerials(urls []string) {
	for _, url := range urls {
		if ctype, err := ContentType(url); err != nil {
			fmt.Printf("error while fetching url %s - %s\n", url, err)
		} else {
			fmt.Printf("Response from %s is %s\n", url, ctype)
		}
	}
}
func siteConcurrent(urls []string) {
	// var wg sync.WaitGroup
	ch := make(chan string)
	for _, url := range urls {
		// wg.Add(1)
		go ContentTypeP(url, ch)
	}
	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}

func ExeSyncSS(urls []string) {
	start := time.Now()
	siteConcurrent(urls)
	fmt.Println("Time Taken: ", time.Since(start))
}
func ExeAsyncSS(urls []string) {
	start := time.Now()
	siteSerials(urls)
	fmt.Println("Time Taken: ", time.Since(start))
}

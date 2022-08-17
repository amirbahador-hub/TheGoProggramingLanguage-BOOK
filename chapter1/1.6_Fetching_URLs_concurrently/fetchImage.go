package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func generateUrl() {
	urls := ""
	for i := 0; i < 10; i++ {
		urls = fmt.Sprintf("%s%d\n", urls+"https://www.picsum.photos/500/800?random=", i)
	}
	urlsByte := []byte(urls)
	err := ioutil.WriteFile("urls.txt", urlsByte, 0644)
	if err != nil {
		return
	}
}

func main() {
	generateUrl()
	start := time.Now()
	ch := make(chan string)
	files, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		return
	}
	for _, url := range strings.Split(string(files), "\n") {
		go fetchUrlWithChannel(url, ch)
	}
	for i := 0; i < len(strings.Split(string(files), "\n")); i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f s elapsed with channel \n", time.Since(start).Seconds())
	start = time.Now()
	for _, url := range strings.Split(string(files), "\n") {
		fmt.Println(fetchUrl(url))
	}
	fmt.Printf("%.2f s elapsed without channel \n", time.Since(start).Seconds())
}

func fetchUrlWithChannel(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s %7d", secs, url, nBytes)
}

func fetchUrl(url string) string {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	nBytes, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return ""
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs %s %7d", secs, url, nBytes)
}

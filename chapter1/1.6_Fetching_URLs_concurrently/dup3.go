package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "strings"
  "time"
  "net/http"
  "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
  // for using in image name
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func main() {
  // 21.26s elapsed
  start := time.Now()
  ch := make(chan string)
  for _, filename := range os.Args[1:] {
    data, _ := ioutil.ReadFile(filename)
    for _, line := range strings.Split(string(data), "\n") {
      go fetch(line, ch)
    }
    for range strings.Split(string(data), "\n")  {
    fmt.Println(<-ch)
  }

  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // sebd ti channel ch
    return
  }

  img, _ := os.Create(randSeq(10)+".jpg")
  defer img.Close()
  nbytes, err := io.Copy(img, resp.Body)

  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


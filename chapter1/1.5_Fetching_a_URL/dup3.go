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
  // 46.94s elapsed
  start := time.Now()
  for _, filename := range os.Args[1:] {
    data, _ := ioutil.ReadFile(filename)
    for _, line := range strings.Split(string(data), "\n") {
      img, _ := os.Create(randSeq(10)+".jpg")
      defer img.Close()

      response, _ := http.Get(line)
      defer response.Body.Close()

      b, _ := io.Copy(img, response.Body)
      fmt.Println("File size: ", b)
    }
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

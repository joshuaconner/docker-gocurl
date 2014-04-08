package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: go run curl.go IP_ADDRESS")
    os.Exit(1)
  }
    url := fmt.Sprintf("http://%s:7001/version", os.Args[1])

    resp, err := http.Get(url)
    if err != nil {
    fmt.Printf("Error getting URL: %s\n", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    fmt.Printf("Error reading body: %s\n", err)
    }
    fmt.Printf("%s\n", body)
}

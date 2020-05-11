package main

import (
 "fmt"
 "log"
 "net/http"
)

var version = 2

func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "Hello 世界... from v", version)
}
func main() {
 http.HandleFunc("/", handler)
 log.Fatal(http.ListenAndServe(":8888", nil))
}

package main

import (
    "fmt"
    "log"
    "net/http"
)

// defining article structure
type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// defining global array of articles to be populated by
// main func and simulate a database
type Articles []Article

// function to handle all requests to our root url
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the homepage!")
    fmt.Println("Endpoint hit: homepage")
}

// function to match URL path hit with a defined function
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":1000", nil))
}

// main function
func main() {
    handleRequests()
}

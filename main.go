package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Movie struct {
    Title string `json:"title"`
    Genre string `json:"genre"`
}

func main() {
    data, err := os.ReadFile("movies.json")
    if err != nil {
        panic(err)
    }

    var movies []Movie
    json.Unmarshal(data, &movies)

    for _, m := range movies {
        fmt.Printf("🎬 %s (%s)\n", m.Title, m.Genre)
    }
}

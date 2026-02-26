
encoding/json

Since most modern apps communicate using JSON (JavaScript Object Notation), the encoding/json package is one of the most important tools in your Go toolkit.

In Go, we move data between Structs (Go's way of defining objects) and JSON strings.

1. The Core Concepts: Marshal & Unmarshal
The encoding/json package uses two main verbs:

Marshal: Converting a Go Struct → JSON string (Sending data out).

Unmarshal: Converting a JSON string → Go Struct (Receiving data in).

2. How to Map Data (Struct Tags)
Since Go uses ProperCase for exported variables and JSON usually uses snake_case or camelCase, we use Struct Tags to tell Go how to translate them.

Go
type User struct {
    FirstName string `json:"first_name"`
    Age       int    `json:"age"`
    IsActive  bool   `json:"is_active"`
}
3. Practical Example: Converting Data
Let's see a program that takes a Go object and turns it into a JSON string you could send to a website.

Go
package main

import (
    "encoding/json"
    "fmt"
)

type Movie struct {
    Title  string  `json:"title"`
    Year   int     `json:"year"`
    Rating float64 `json:"rating"`
}

func main() {
    // 1. Create a Go object
    myMovie := Movie{Title: "Inception", Year: 2010, Rating: 8.8}

    // 2. Marshal: Convert to JSON (returns bytes)
    jsonData, _ := json.Marshal(myMovie)

    // 3. Print the result as a string
    fmt.Println(string(jsonData))
    // Output: {"title":"Inception","year":2010,"rating":8.8}
}
4. Reading JSON (Unmarshaling)
If you receive JSON from an API, you do the reverse:

Go
rawJSON := `{"title":"The Matrix","year":1999,"rating":8.7}`
var incomingMovie Movie

// We use & (a pointer) so the function can "fill up" our variable
json.Unmarshal([]byte(rawJSON), &incomingMovie)

fmt.Println(incomingMovie.Title) // Output: The Matrix
Pro Tips for encoding/json
Exported Fields: Only fields starting with a Capital Letter (like Title) can be converted to JSON. If you use title, the json package will ignore it!

Omitempty: You can add ,omitempty to a tag (e.g., json:"id,omitempty") so that if the value is empty/zero, it won't appear in the final JSON at all.

JSON to Go Tool: If you have a massive, complex JSON blob and don't want to write the Struct by hand, use a tool like https://mholt.github.io/json-to-go/ to generate the code for you instantly.



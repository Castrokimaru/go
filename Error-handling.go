1. The Go Error Pattern
When a function might fail (like reading a file or parsing JSON), it returns two things: the result and an error. If the error is nil, everything went fine.

Go
data, err := someFunction()
if err != nil {
    // Handle the error (log it, return it, or stop)
    fmt.Println("Something went wrong:", err)
    return 
}
// If we are here, data is safe to use!
2. Why do it this way?
No Surprises: You are forced to deal with the "bad path" immediately.

Control: You decide exactly how to recover from a specific error.

Readability: You can follow the "Happy Path" of the code down the left side of the screen.

3. Building the JSON API (With Error Handling)
Now, let's combine everything. We will build an API that takes a "Book" request and returns a JSON response, but it will handle errors if the user sends us "garbage" data.

Go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Only allow POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var newBook Book
    // 2. Try to decode the JSON. If it fails, handle the ERROR.
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, "Invalid JSON data", http.StatusBadRequest)
        return
    }

    // 3. If successful, send a custom response back
    fmt.Printf("Received book: %s by %s\n", newBook.Title, newBook.Author)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
    http.HandleFunc("/add-book", bookHandler)
    fmt.Println("API running on :8080...")
    http.ListenAndServe(":8080", nil)
}
What’s happening in that code?
json.NewDecoder: This is a memory-efficient way to read JSON directly from the request body.

http.Error: A helper function that sets the correct status code (like 400 or 405) and writes the error message for you.

The Pointer (&newBook): Notice the & symbol. We pass the address of our struct so the Decode function can reach into our memory and fill it with data.

Practice Tip
If you run this locally, you can test it using a tool like Postman or a simple curl command in your terminal:
curl -X POST -d '{"title":"The Go Programming Language", "author":"Alan Donovan"}' http://localhost:8080/add-book
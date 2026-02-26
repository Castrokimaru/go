 let’s look at the net/http package. This is where Go truly shines.
 In many other languages, you need a massive "framework" (like Spring or Django) to build a web server. 
 In Go, it’s built right into the standard library.

 //Building a Web Server in 10 Lines
-This program starts a local server that says "Hello" to anyone who visits it in a browser.

package main

import (
    "fmt"
    "net/http" // The magic web package
)

func main() {
    // 1. Define a "Route" (what happens when someone visits "/")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my Go Web Server!")
    })

    fmt.Println("Server starting on http://localhost:8080...")

    // 2. Start the server on port 8080
    http.ListenAndServe(":8080", nil)
}


How the net/http Package Works
The net/http package follows a simple Request/Response pattern. When a user visits your URL, Go creates two objects for you:

http.Request (r): This contains everything the user sent you (their IP, their browser type, any data they submitted).

http.ResponseWriter (w): This is your "pen." Whatever you write to w gets sent back to the user's browser.


Why this is a Big DealPerformance: This built-in server is production-ready.
 It can handle thousands of concurrent users right out of the box because Go automatically gives every single request its own "Goroutine" (a lightweight thread).
 No Dependencies: You don't have to download 500MB of libraries to get a website running. 
 It’s all in that one net/http package.
 //
 // Common net/http Functions 
 // Function                      Purpose
http.HandleFunc            Maps a URL path (like /login) to a function.
http.ListenAndServe        Keeps the program running and listening for visitors.
http.Get                   Acts like a browser to go "fetch" data from another website.
http.Post                  Sends data (like a form) to another server.
package main

import (
    "fmt"
    "log"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    fmt.Fprintf(w, "POST request successful")

    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        // Doesn't close the request, the caller must avoid further writings to the response after calling this.
        http.Error(w, "Method is not supported", http.StatusMethodNotAllowed) 
    }

    // Formats a string and write it to a io w.
    fmt.Fprintf(w, "hello!")
}

func main(){
    // http.Dir here is a type casting to cast the type string to http.Dir, which implements the method
    // FileSystem required by the http.FileServer method.
    fileServer := http.FileServer(http.Dir("./static/")) 
    http.Handle("/", fileServer) // Handle expects an Handler interface witch implements the ServeHTTP method.
    http.HandleFunc("/form", formHandler) // HandleFunc expects a function with a ResponseWriter and a Request.
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err:= http.ListenAndServe(":8080", nil); err != nil {
        // Writes to stderr with date and time, calls os.Exit(1) after.
        log.Fatal(err)
    }
}

package main

import (
	"fmt"
	"log"
	"net/http"
    // "goNoFrameWork/routes"
    "goNoFrameWork/render"
) 

func main() {
    // This is a server that will return hello world served on /hello
    var port = ":8000"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        err := render.RenderTemplate(w, r, "\\index.html")
        if err != nil {
            fmt.Printf("ERROR cannot render")
        }
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        err := render.RenderTemplate(w, r, "\\about.html")
        if err != nil {
            fmt.Printf("ERROR cannot render")
        }
    })

    fmt.Printf("Serving at http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}


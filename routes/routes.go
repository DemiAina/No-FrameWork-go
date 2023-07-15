package routes

import (
    "net/http"
    "io"
)

func HelloHandler  (w http.ResponseWriter, req *http.Request){
    io.WriteString(w, "Hello World!\n")
}


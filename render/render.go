package render

import (
    "fmt"
    "html/template"
    "net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request ,  html string) error{
    parseTemplate, err := template.ParseFiles("./html" + html)
    fmt.Println("\\html" + html)
    if err != nil{
        fmt.Printf("ERROR parsing file")
    }
    err = parseTemplate.Execute(w, nil)
    if err != nil{
        fmt.Printf("ERROR excuting")
    }
    
    return nil
}

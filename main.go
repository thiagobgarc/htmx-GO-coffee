package main 

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

type Coffee struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Price float64 `json:"price"`
    Description string `json:"description"`
}

func main() {
    fmt.Println("Coffee Menu")

    h1 := func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("index.html"))
        coffee := map[string][]Coffee{
            "Coffee": {
                {ID: "1", Title: "Small", Price: 1.99},
                {ID: "2", Title: "Medium", Price: 2.99},
                {ID: "3", Title: "Large", Price: 3.99},
            },
        }
        tmpl.Execute(w, coffee)
    }

    h2 := func(w http.ResponseWriter, r *http.Request) {
        title := r.PostFormValue("title")
        description := r.PostFormValue("description")
        fmt.Println(title)
        fmt.Println(description)

    }

    http.HandleFunc("/coffee", h1)
    http.HandleFunc("/add-coffee/", h2)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

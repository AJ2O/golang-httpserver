package main

import (
	// Provides templating language for HTML templates
	"text/template"
	// Includes automatic escaping of data before displaying it
	// in the browser
	// ex. XSS attacks
	"log"
	"net/http"
)

// Todo - a to-do item.
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData - a page of to-do items.
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// Control structures
	// {{/* comment */}}                Defines a comment
	// {{.}}                            Renders the root element
	// {{.Title}}                       Renders the "Title"-field in a nested element
	// {{if .Done}} {{else}} {{end}}    Defines an if-Statement
	// {{range .Todos}} {{.}} {{end}}   Loops over all "Todos" and renders each
	//                                  using {{.}}
	// {{block "content" .}} {{end}}    Defines a block with the name "content"

	// Templates can also be parsed from string or from files
	// tmpl, err := template.ParseFiles("layout.html")
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		log.Fatal(err)
	}

	// Writing out the template
	// func(w http.ResponseWriter, r *http.Request) {
	//     tmpl.Execute(w, "data goes here")
	// }

	// For template rendering, the data passed can be any kind of Go data structure
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}

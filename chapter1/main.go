package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler represents a single template
type TemplateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// NewTemplateHandler ...
func NewTemplateHandler(mytemplate string) *TemplateHandler {

	return &TemplateHandler{filename: mytemplate}

}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	chattemplate := &TemplateHandler{filename: "chat.html"}

	//http.Handle("/", NewTemplateHandler("chat.html"))

	//	http.HandleFunc("/")
	http.Handle("/", chattemplate)

	r := newRoom()
	r.sayHello()
	//fmt.Println(r.sayHello())

	http.Handle("/room", r)
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	templates := parseTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates["index.html"].Execute(w, struct{ Title string }{Title: "Index"})
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		templates["about.html"].Execute(w, struct{ Title string }{Title: "About"})
	})

	http.ListenAndServe(":3000", nil)
}

func parseTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)

	layout, _ := template.ParseFiles("templates/_layout.html")

	dir, _ := os.Open("templates/blocks")
	defer dir.Close()

	fis, _ := dir.Readdir(-1)

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		f, _ := os.Open("templates/blocks/" + fi.Name())
		content, _ := ioutil.ReadAll(f)
		f.Close()
		tmpl, _ := layout.Clone()
		tmpl.Parse(string(content))
		result[fi.Name()] = tmpl
	}

	return result
}

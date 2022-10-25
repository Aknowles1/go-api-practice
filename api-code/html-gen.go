package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("./html/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Println("filename to save to: " + filename)
	return os.WriteFile("./html/"+filename, p.Body, 0600)
}

func (a *application) homePage(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	fmt.Println("page title: " + title)
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func (a *application) editHome(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := loadPage("templatedhome")
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("./html/edit.html")
	t.Execute(w, p)
}

func (a *application) saveHome(w http.ResponseWriter, r *http.Request) {
	title := "templatedhome"
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		fmt.Printf("error: %v" + err.Error())
	}
	http.Redirect(w, r, "/templatedhome", http.StatusFound)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	fileName := p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r * http.Request){
	titel := r.URL.Path[len("/edit/")]
	p, _ := laodPage(title)
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", p.Title, p.Body)
}

func main(){

	fmt.Println("server is up")


	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//p1 := &Page{Title: "test", Body: []byte("This is sample Page.")}
	//saveErr := p1.save()
	//if saveE
	//rr != nil {
	//	fmt.Println(saveErr)
	//}
	//
	//
	//p2, loadErr := loadPage(p1.Title)
	//if loadErr != nil {
	//	fmt.Println(loadErr)
	//}
	//fmt.Println(string(p2.Body))
}
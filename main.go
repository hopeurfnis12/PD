package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name: %s\n Age: %d\nMoney: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 2.0}
	//bob.money = -23
	//bob.setNewName("asd")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Main text</b>")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}

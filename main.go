package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Subjects struct {
	Id      uint16
	Subject string
	Sum     float32
}

var subjs = []Subjects{}

/* ///////////////// HOME ///////////////// */
func home_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `subjects`")
	if err != nil {
		panic(err)
	}

	subjs = []Subjects{}
	for res.Next() {
		var subj Subjects
		err = res.Scan(&subj.Id, &subj.Subject, &subj.Sum)
		if err != nil {
			panic(err)
		}

		subjs = append(subjs, subj)
	}

	t.ExecuteTemplate(w, "home_page", subjs)
}

/* ///////////////// SUBJECTS ///////////////// */
func subjects_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/subjects_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "subjects_page", nil)
}

/* ///////////////// ADD ///////////////// */
func add_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/add_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "add_page", nil)
}

/* ///////////////// SAVE ///////////////// */
func save(w http.ResponseWriter, r *http.Request) {
	subj := r.FormValue("subj")
	sum := r.FormValue("sum")

	if subj == "" || sum == "" {
		http.Redirect(w, r, "/add#error", http.StatusSeeOther)
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diary")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `subjects` (`subject`, `sum`) VALUES('%s', '%s')", subj, sum))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleRequest() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/subjects/", subjects_page)
	http.HandleFunc("/add/", add_page)
	http.HandleFunc("/save/", save)
	http.ListenAndServe(":7272", nil)
}

// to start the server write in terminal > go run main.go
func main() {
	fmt.Println("Server is started (to stop press: 'ctrl' + 'c')")
	handleRequest()
}

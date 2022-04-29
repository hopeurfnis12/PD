package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Subjects struct {
	Id      uint16
	Subject string
	Sum     float32
}

type Todo_list struct {
	Id         uint16
	Subject_id uint16
	Todo       string
	Score      float32
}

var subjs = []Subjects{}
var todo_list = []Todo_list{}

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
		http.Redirect(w, r, "/add?error", http.StatusSeeOther)
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

/* ///////////////// SUBJECT ///////////////// */
func subject_show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/subject_show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `todo_list` WHERE `subject_id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	// m := map[string]interface{}{}
	// m["todo_list"] = []Todo_list{}
	// m["title"], err = db.Query(fmt.Sprintf("SELECT `subject` FROM `subjects` WHERE `id` = '%s'", vars["id"]))
	// if err != nil {
	// 	panic(err)
	// }

	todo_list = []Todo_list{}
	for res.Next() {
		var todo Todo_list
		err = res.Scan(&todo.Id, &todo.Subject_id, &todo.Todo, &todo.Score)
		if err != nil {
			panic(err)
		}

		todo_list = append(todo_list, todo)
	}

	t.ExecuteTemplate(w, "subject_show", todo_list)
}

/* ///////////////// handle ///////////////// */
func handleRequest() {
	rtr := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	rtr.HandleFunc("/", home_page).Methods("GET")
	rtr.HandleFunc("/subjects/", subjects_page).Methods("GET")
	rtr.HandleFunc("/add/", add_page).Methods("GET")
	rtr.HandleFunc("/save/", save).Methods("POST")
	rtr.HandleFunc("/subject/{id:[0-9]+}", subject_show).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":7272", nil)
}

// to start the server write in terminal > go run main.go
func main() {
	fmt.Println("Server is started (to stop press: 'ctrl' + 'c')")
	handleRequest()
}

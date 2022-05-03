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
	Count   int16
}

type TodoList struct {
	Id        uint16
	SubjectID uint16
	Todo      string
	Do        int16
}

var user = "root"
var password = ""
var subjs = []Subjects{}
var todo_list = []TodoList{}

/* ///////////////// HOME ///////////////// */
func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "HomePage", subjs)
}

/* ///////////////// SUBJECTS ///////////////// */
func subjects_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/subjects_page.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `subjects` ORDER BY `Id` DESC")
	if err != nil {
		panic(err)
	}

	subjs = []Subjects{}
	for res.Next() {
		var subj Subjects
		err = res.Scan(&subj.Id, &subj.Subject)
		if err != nil {
			panic(err)
		}

		er := db.QueryRow("SELECT COUNT(*) FROM `todo_list` WHERE `do` = 0 AND `subject_id` = ?", subj.Id).Scan(&subj.Count)
		if er != nil {
			panic(er)
		}

		subjs = append(subjs, subj)
	}

	t.ExecuteTemplate(w, "subjects_page", subjs)
}

/* ///////////////// SAVE ///////////////// */
func save(w http.ResponseWriter, r *http.Request) {
	subj := r.FormValue("subj")

	if subj == "" {
		http.Redirect(w, r, "/?error", http.StatusSeeOther)
	} else {
		db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `subjects` (`subject`) VALUES('%s')", subj))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/subjects/", http.StatusSeeOther)
	}
}

/* ///////////////// EDIT_SUBJ ///////////////// */
func edit_subj(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjed := r.FormValue("subj-edit")

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ed, err := db.Query(fmt.Sprintf("UPDATE `subjects` SET `subject`='%s' WHERE `id` = '%s'", subjed, vars["id"]))
	if err != nil {
		panic(err)
	}
	defer ed.Close()

	http.Redirect(w, r, "/subjects/", http.StatusSeeOther)
}

/* ///////////////// DEL_SUBJ ///////////////// */
func del_subj(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	upd, err := db.Query(fmt.Sprintf("DELETE FROM `subjects` WHERE `id`='%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	defer upd.Close()

	http.Redirect(w, r, "/subjects/", http.StatusSeeOther)
}

/* ///////////////// SUBJECT ///////////////// */
func subject_show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/subject_show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `todo_list` WHERE `subject_id` = '%s' ORDER BY `Id` DESC", vars["id"]))
	if err != nil {
		panic(err)
	}

	todo_list = []TodoList{}
	for res.Next() {
		var todo TodoList
		err = res.Scan(&todo.Id, &todo.SubjectID, &todo.Todo, &todo.Do)
		if err != nil {
			panic(err)
		}

		todo_list = append(todo_list, todo)
	}

	t.ExecuteTemplate(w, "subject_show", todo_list)
}

/* ///////////////// SAVE_TASK ///////////////// */
func save_task(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	task := r.FormValue("task")

	if task == "" {
		http.Redirect(w, r, "/?error", http.StatusSeeOther)
	} else {
		db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `todo_list` (`subject_id`, `todo`, `do`) VALUES('%s', '%s', 0)", vars["id_subj"], task))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/subject/"+vars["id_subj"], http.StatusSeeOther)
	}
}

/* ///////////////// DO ///////////////// */
func do(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var do_make = 0

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if vars["do"] == "1" {
		do_make = 0
	} else {
		do_make = 1
	}
	upd, err := db.Query(fmt.Sprintf("UPDATE `todo_list` SET `do` = %d WHERE `id` = %s", do_make, vars["id_task"]))
	if err != nil {
		panic(err)
	}
	defer upd.Close()

	http.Redirect(w, r, "/subject/"+vars["id_subj"], http.StatusSeeOther)
}

/* ///////////////// EDIT ///////////////// */
func edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tasked := r.FormValue("task-edit")

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ed, err := db.Query(fmt.Sprintf("UPDATE `todo_list` SET `todo`='%s' WHERE `id` = '%s'", tasked, vars["id_task"]))
	if err != nil {
		panic(err)
	}
	defer ed.Close()

	http.Redirect(w, r, "/subject/"+vars["id_subj"], http.StatusSeeOther)
}

/* ///////////////// DEL ///////////////// */
func del(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/diary")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	del, err := db.Query(fmt.Sprintf("DELETE FROM `todo_list` WHERE `id`='%s'", vars["id_task"]))
	if err != nil {
		panic(err)
	}
	defer del.Close()

	http.Redirect(w, r, "/subject/"+vars["id_subj"], http.StatusSeeOther)
}

/* ///////////////// handle ///////////////// */
func handleRequest() {
	rtr := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	rtr.HandleFunc("/", HomePage).Methods("GET")

	rtr.HandleFunc("/subjects/", subjects_page).Methods("GET")
	rtr.HandleFunc("/save/", save).Methods("POST")
	rtr.HandleFunc("/edit_subj/{id:[0-9]+}", edit_subj).Methods("POST", "GET")
	rtr.HandleFunc("/del_subj/{id:[0-9]+}", del_subj).Methods("POST", "GET")

	rtr.HandleFunc("/subject/{id:[0-9]+}", subject_show).Methods("GET")
	rtr.HandleFunc("/save_task/{id_subj:[0-9]+}", save_task).Methods("POST")
	rtr.HandleFunc("/do/{id_subj:[0-9]+}/{id_task:[0-9]+}/{do:[0-1]}", do).Methods("POST", "GET")
	rtr.HandleFunc("/edit/{id_subj:[0-9]+}/{id_task:[0-9]+}", edit).Methods("POST", "GET")
	rtr.HandleFunc("/del/{id_subj:[0-9]+}/{id_task:[0-9]+}", del).Methods("POST", "GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":7272", nil)
}

// to start the server write in terminal > go run main.go
func main() {
	fmt.Println("Server is started (to stop press: 'ctrl' + 'c')")
	handleRequest()
}

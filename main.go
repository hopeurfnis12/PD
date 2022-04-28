package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name       string
	Age        uint16
	Money      int16
	Avg_grades float64
	Hobbies    []string
}

//type Subjects struct {
// Subject string  `json:"subject"`
// Score   float64 `json:"score"`
//}

// func (u *User) getAllInfo() string {
// 	return fmt.Sprintf("User name: %s\n Age: %d\nMoney: %d", u.Name, u.Age, u.Money)
// }

// func (u *User) setNewName(newName string) {
// 	u.Name = newName
// }

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, []string{"Info", "RED", "Coding"}}
	//bob.money = -23
	//bob.setNewName("asd")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl.ExecuteTemplate(w, "home_page.html", bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Main text</b>")
}

func handleRequest() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":7272", nil)
}

// to start the server write in terminal > go run main.go
func main() {
	fmt.Println("Server is started (to stop press: 'ctrl' + 'c')")
	handleRequest()
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diary")
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// // insert, err := db.Query("INSERT INTO `subjects` (`subject`, `score`) VALUES('Математика', 5)")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // defer insert.Close()

	// res, err := db.Query("SELECT `subject`, `score` FROM `subjects`")
	// if err != nil {
	// 	panic(err)
	// }

	// for res.Next() {
	// 	var subjects Subjects
	// 	err = res.Scan(&subjects.Subject, &subjects.Score)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(fmt.Sprintf("Subject: %s, Score: %f", subjects.Subject, subjects.Score))
	// }

}

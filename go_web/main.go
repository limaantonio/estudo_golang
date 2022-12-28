package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"text/template"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var connStr = "user=postgres password=root dbname=go_course sslmode=disable"
var db, err = sql.Open("postgres", connStr)

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("/static"))))
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func GetPostById(id string) Post {
	row := db.QueryRow("select * from posts where id=?", id)
	post := Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return post
}

func ListPosts() []Post {
	rows, err := db.Query("select * from posts")
	check(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)

	}

	return items
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/layout.html", "tamplates/list.html"))
	if err := t.ExecuteTemplate(w, "layout.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	t := template.Must(template.ParseFiles("templates/layout.html", "tamplates/view.html"))
	if err := t.ExecuteTemplate(w, "layout.html", GetPostById(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 	sqlStatement := `Insert into posts(title, body)
// VALUES ($1, $2)`
// 	_, err = db.Exec(sqlStatement, "Unamed Post", "Our content")
// 	if err != nil {
// 		panic(err)
// 	}

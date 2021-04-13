package main

import (
	"html/template"
	"net/http"
)

func studentList(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	//preparing data for sending to frontend
	if session.Values["isLoggedIn"] == nil {
		session.Values["isLoggedIn"] = false
		session.Values["username"] = ""
	}
	data := struct {
		Title      string
		IsLoggedIn bool
		Username   string
		Lists      []map[string]interface{}
	}{
		Title:      "Student List | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
		Lists:      getStudentList("company::1"),
	}

	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl, err = tmpl.ParseFiles("wpage/student_list.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
}

package main

import (
	"html/template"
	"net/http"
	"strings"
)

func studentList(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	//** process starts: preparing data for sending to frontend **//
	if session.Values["isLoggedIn"] == nil {
		session.Values["isLoggedIn"] = false
		session.Values["username"] = ""
	}

	// using struct literal
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
	//** process ends: preparing data for sending to frontend **//

	//** process starts: executing template **//
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl, err = tmpl.ParseFiles("wpage/student_list.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
	//** process ends: executing template **//
}

func studentSearch(w http.ResponseWriter, r *http.Request) {
	//** process starts: getting form data **//
	formData := make(map[string]string)
	formData["aid"] = strings.TrimSpace(r.FormValue("student-id"))
	formData["account_name"] = strings.ToLower(strings.TrimSpace(r.FormValue("name")))
	formData["email"] = strings.TrimSpace(r.FormValue("email"))
	formData["mobile"] = strings.TrimSpace(r.FormValue("phone"))
	//** process ends: getting form data **//

	session, err := store.Get(r, "mysession")
	checkErr(err)
	//** process starts: preparing data for sending to frontend **//
	if session.Values["isLoggedIn"] == nil {
		session.Values["isLoggedIn"] = false
		session.Values["username"] = ""
	}

	// using struct literal
	data := struct {
		Title      string
		IsLoggedIn bool
		Username   string
		Lists      []map[string]interface{}
	}{
		Title:      "Student List | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
		Lists:      doSearch(formData, "company::1"),
	}
	//** process ends: preparing data for sending to frontend **//

	//** process starts: executing template **//
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl, err = tmpl.ParseFiles("wpage/student_list.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
	//** process ends: executing template **//
}

package main

import (
	"html/template"
	"net/http"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if session.Values["isLoggedIn"] == true {
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
		}{
			Title:      "Dashboard | MASTER-ACADEMY",
			IsLoggedIn: session.Values["isLoggedIn"].(bool),
			Username:   session.Values["username"].(string),
		}
		//** process ends: preparing data for sending to frontend **//

		//** process starts: executing template **//
		tmpl, err := template.ParseFiles("template/index.gohtml")
		checkErr(err)
		tmpl, err = tmpl.ParseFiles("wpage/dashboard.gohtml")
		checkErr(err)
		tmpl.Execute(w, data)
		//** process ends: executing template **//
	} else { //if not logged in then redirecting to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

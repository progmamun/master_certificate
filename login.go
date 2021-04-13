package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if r.Method != "POST" {
		if session.Values["isLoggedIn"] == true { // if already logged in, then redirecting to homepage
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
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
				Title:      "Login | MASTER-ACADEMY",
				IsLoggedIn: session.Values["isLoggedIn"].(bool),
				Username:   session.Values["username"].(string),
			}
			//** process ends: preparing data for sending to frontend **//

			//** process starts: executing template **//
			tmpl, err := template.ParseFiles("template/index.gohtml")
			checkErr(err)
			tmpl, err = tmpl.ParseFiles("wpage/login.gohtml")
			checkErr(err)
			tmpl.Execute(w, data)
			//** process ends: executing template **//
		}
	} else {
		//** process starts: getting form data **//
		username := strings.TrimSpace(r.FormValue("username"))
		password := r.FormValue("password")
		//** process ends: getting form data **//

		check := doLogin(username, password) // calling a function which will try to login

		if check == "Done" { // if login done
			//** process ends: storing session values **//
			session.Values["username"] = username
			session.Values["password"] = password
			session.Values["isLoggedIn"] = true
			session.Save(r, w)
			//** process ends: storing session values **//

			fmt.Fprintln(w, "Login Done")
		} else if check == "username" {
			fmt.Fprintln(w, "Username not found!")
		} else {
			fmt.Fprintln(w, "Wrong password!")
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	//** process starts: deleting session values **//
	session.Values["username"] = ""
	session.Values["password"] = ""
	session.Values["isLoggedIn"] = false

	session.Options.MaxAge = -1 //cookies will be deleted immediately.
	session.Save(r, w)
	//** process ends: deleting session values **//

	http.Redirect(w, r, "/", http.StatusSeeOther) //after logout, redirecting to homepage
}

func forgotpassword(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	//preparing data for sending frontend
	if session.Values["isLoggedIn"] == nil {
		session.Values["isLoggedIn"] = false
		session.Values["username"] = ""
	}
	data := struct {
		Title      string
		IsLoggedIn bool
		Username   string
	}{
		Title:      "Request for password reset | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
	}

	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl, err = tmpl.ParseFiles("wpage/forgot_password.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
}

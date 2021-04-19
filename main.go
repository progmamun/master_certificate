package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	//Just a message for ensuring the local server is running
	fmt.Println("Local server is listening on port 9001...")

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/forgot_password", forgotpassword)
	http.HandleFunc("/dashboard", dashboard)

	http.HandleFunc("/student_list", studentList)
	http.HandleFunc("/student_search", studentSearch)

	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.HandleFunc("/api/", api)
	http.HandleFunc("/cert", cert)

	//serving file from server to client
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	//localhost running on port 9001
	http.ListenAndServe(":9001", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
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
	}{
		Title:      "Certificate Generator | MASTER-ACADEMY",
		IsLoggedIn: session.Values["isLoggedIn"].(bool),
		Username:   session.Values["username"].(string),
	}
	//** process ends: preparing data for sending to frontend **//

	//** process starts: executing template **//
	tmpl, err := template.ParseFiles("template/index.gohtml")
	checkErr(err)
	tmpl.Execute(w, data)
	//** process ends: executing template **//
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contact page")
}

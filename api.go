package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func api(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/api/update" {
		firstName := strings.TrimSpace(r.FormValue("firstName"))
		lastName := strings.TrimSpace(r.FormValue("lastName"))
		username := strings.TrimSpace(r.FormValue("username"))
		mobile := strings.TrimSpace(r.FormValue("mobile"))
		city := strings.TrimSpace(r.FormValue("city"))
		status := strings.TrimSpace(r.FormValue("status"))
		sts, _ := strconv.Atoi(status)

		res := updateStudentInfo(firstName, lastName, username, mobile, city, sts)
		fmt.Fprintln(w, res)
		return
	}

	if path == "/api/register" {
		formData := make(map[string]string)
		formData["firstName"] = strings.TrimSpace(r.FormValue("firstName"))
		formData["lastName"] = strings.TrimSpace(r.FormValue("lastName"))
		formData["username"] = strings.TrimSpace(r.FormValue("username"))
		formData["email"] = strings.TrimSpace(r.FormValue("email"))
		formData["password"] = r.FormValue("password")
		formData["remarks"] = strings.TrimSpace(r.FormValue("remarks"))

		doRegistration(formData, w, r)
		return
	}

	index := strings.Index(path, "/api/student_info/")
	if index != -1 { //url is like this "/api/student_info/....."
		processStudentInfo(path, w)
		return
	}
	//errorPage(w, http.StatusBadRequest) //http.StatusBadRequest = 400
}

// processStudentInfo() function returns a single student information to frontend
func processStudentInfo(path string, w http.ResponseWriter) {
	username := strings.TrimPrefix(path, "/api/student_info/")
	info := getSingleStudent(username)

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(info[0])
	w.Write(b)
}

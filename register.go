package main

import (
	"html/template"
	"net/http"
	"strings"
)

func register(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	checkErr(err)

	if r.Method != "POST" {
		if session.Values["isLoggedIn"] == true {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			//preparing data for sending to frontend
			if session.Values["isLoggedIn"] == nil {
				session.Values["isLoggedIn"] = false
				session.Values["username"] = ""
			}
			data := struct {
				Title      string
				IsLoggedIn bool
				Username   string
			}{
				Title:      "Register | MASTER-ACADEMY",
				IsLoggedIn: session.Values["isLoggedIn"].(bool),
				Username:   session.Values["username"].(string),
			}

			tmpl, err := template.ParseFiles("template/index.gohtml")
			checkErr(err)
			tmpl, err = tmpl.ParseFiles("wpage/register.gohtml")
			checkErr(err)
			tmpl.Execute(w, data)
		}
	} else {
		//getting form data
		formData := make(map[string]string)
		formData["firstName"] = strings.TrimSpace(r.FormValue("firstName"))
		formData["lastName"] = strings.TrimSpace(r.FormValue("lastName"))
		formData["username"] = strings.TrimSpace(r.FormValue("username"))
		formData["email"] = strings.TrimSpace(r.FormValue("email"))
		formData["password"] = r.FormValue("password")

		doRegistration(formData, w, r) //if not exsist, then proceed to registration
	}
}

// func generateToken() string {
// 	b := make([]byte, 16)
// 	rand.Read(b)

// 	hasher := md5.New()
// 	hasher.Write(b)

// 	return hex.EncodeToString(hasher.Sum(nil))
// }

// func sendMail(email, username, link string) {
// 	auth := smtp.PlainAuth("", "fakenahid@gmail.com", "hqfumidtzssgmdzr", "smtp.gmail.com")
// 	to := []string{email}

// 	//var msg []byte
// 	var body bytes.Buffer

// 	subject := "Master-Academy Account Verification"
// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	msg := fmt.Sprintf("From: Master-Academy \nSubject: %s \nTo:%s \n%s\n\n", subject, email, mimeHeaders)
// 	body.Write([]byte(msg))

// 	data := struct {
// 		Username, Link string
// 	}{
// 		Username: username,
// 		Link:     link,
// 	}
// 	tmpl, err := template.ParseFiles("template/mail.gohtml")
// 	checkErr(err)
// 	tmpl.Execute(&body, data)

// 	err = smtp.SendMail("smtp.gmail.com:587", auth, "", to, body.Bytes())
// 	checkErr(err)
// }

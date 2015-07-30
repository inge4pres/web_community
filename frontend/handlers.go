package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"web_community/backend"
	"web_community/community"
)

var templateDir = "../templates/"
var templates = template.Must(template.ParseGlob(templateDir + "*.tmpl"))

var scookie = securecookie.New(securecookie.GenerateRandomKey(1024), securecookie.GenerateRandomKey(4096))
var store = new(sessions.CookieStore)

type webRes struct {
	Status int
	Err    error
	Title  string
	User   *model.User
}

func indexHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w := new(webRes)
	w.Status = http.StatusOK
	w.Title = "Web Community"

	if session := getSession(req); session.IsNew {
		w.User = nil
	} else {
		w.User.Id, w.Err = sessiondao.GetUserIdBySession(session.Values["session_id"].(string))
	}

	w.Err = templates.ExecuteTemplate(res, "index.tmpl", w)
	if w.Err != nil {
		http.Error(res, w.Err.Error(), http.StatusInternalServerError)
		return
	}
}

func signInHandler(res http.ResponseWriter, req *http.Request) {
	return
}

func signUpHandler(res http.ResponseWriter, req *http.Request) {
	return
}

func getSession(req *http.Request) *sessions.Session {
	session, _ := store.Get(req, "wc_user")
	return session
}

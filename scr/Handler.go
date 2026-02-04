package jeux

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func ConnexionHandler(w http.ResponseWriter, r *http.Request, objet *Video) {
	*objet = Video{
		Pseudo: r.FormValue("pseudo"),
	}
	cookie := &http.Cookie{
		Name:  "cookie",
		Value: objet.Pseudo,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/affichage", http.StatusFound)
}

func AffichageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Page/affichage.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

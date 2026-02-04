package jeux

import (
	"encoding/json"
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

func ConnexionHandler(w http.ResponseWriter, r *http.Request, game *Game, pseudo *User) {
	pseudo.Pseudo = r.FormValue("pseudo")
	AddTableUser(&pseudo.Pseudo)

	resp, err := http.Get("https://www.freetogame.com/api/games")
	if err != nil {
		log.Fatal(err)
	}

	var games []Game

	json.NewDecoder(resp.Body).Decode(&games)
	println(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// AddTableGamesFavoris(&games, &pseudo.Id)

	cookie := &http.Cookie{
		Name:  "cookie",
		Value: pseudo.Pseudo,
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

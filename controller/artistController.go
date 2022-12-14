package controller

import (
	"balldontlie/service"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	t, err := template.ParseFiles("templates/base.html", "templates/404.html")

	if err != nil {
		handle500(w, err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		handle500(w, err)
		return
	}
}

func handle400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)

	t, err := template.ParseFiles("templates/base.html", "templates/400.html")

	if err != nil {
		handle500(w, err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		handle500(w, err)
		return
	}
}

func handle500(w http.ResponseWriter, err error) {
	w.WriteHeader(500)

	t, other := template.ParseFiles("templates/base.html", "templates/500.html")

	if other != nil {
		w.Write([]byte("Something went wrong\nError 500\n" + other.Error()))
		return
	}

	fmt.Println(err)
	t.Execute(w, err)
}

func Get(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/artist/"):]
	if len(id) == 0 {
		w.Header().Set("Content-Type", "application/json")
		if err := getPlayers(w); err != nil {
			handle500(w, err)
			return
		}
	} else {
		idPlayer, err := strconv.Atoi(id)
		if err != nil {
			handle404(w, r)
			return
		}

		artist, err := service.GetPlayerById(idPlayer)

		t, err := template.ParseFiles("templates/base.html", "templates/artist.html")
		if err != nil {
			handle400(w, r)
			return
		}

		err = t.Execute(w, artist)
		if err != nil {
			handle400(w, r)
			return
		}
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, err := template.ParseFiles("templates/base.html", "templates/home.html")

		if err != nil {
			handle500(w, err)

			return
		}
		artists, err := service.Get()
		if err != nil {
			handle500(w, err)

			return
		}
		err = t.Execute(w, artists)
		if err != nil {
			handle500(w, err)

			return
		}
	} else {
		handle404(w, r)
		return
	}
}

func getPlayers(w http.ResponseWriter) error {
	players, err := service.Get()
	if err != nil {
		return err
	}

	playersJson, err := json.Marshal(players)
	if err != nil {
		return err
	}

	_, err = w.Write(playersJson)
	if err != nil {
		return err
	}

	return nil
}

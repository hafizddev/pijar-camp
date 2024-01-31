package provinsicontroller

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"pijar_camp/entities"
	"pijar_camp/models/provinsimodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/provinsi/index.html")
	if err != nil {
		panic(err)
	}

	provinsi := provinsimodel.GetAll()
	data := map[string]any{
		"provinsi": provinsi,
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/provinsi/create.html")
	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodGet {
		temp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		provinsi := entities.Provinsi{
			Name:       r.FormValue("name"),
			CreatedAt:  time.Now(),
			UpddatedAt: time.Now(),
		}
		if ok := provinsimodel.Create(provinsi); !ok {
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/provinsi", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/provinsi/edit.html")
	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodGet {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		provinsi := provinsimodel.Detail(id)
		data := map[string]any{
			"provinsi": provinsi,
		}

		temp.Execute(w, data)
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		provinsi := entities.Provinsi{
			Name:       r.FormValue("name"),
			UpddatedAt: time.Now(),
		}

		if ok := provinsimodel.Update(id, provinsi); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/provinsi", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := provinsimodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/provinsi", http.StatusSeeOther)
}

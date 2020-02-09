package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"mattmeinzer.com/plants/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	p, err := app.plants.Top()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Plants: p,
	})
}

func (app *application) showPlant(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	plant, err := app.plants.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Plant: plant,
	})
}

func (app *application) createPlantForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create plant form"))
}

func (app *application) createPlant(w http.ResponseWriter, r *http.Request) {
	name := "Timothy" // FIXME - actually use the params

	id, err := app.plants.Insert(name)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/plant?id=%d", id), http.StatusSeeOther)
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createPlant(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	name := r.PostForm.Get("name")

	errors := make(map[string]string)

	if strings.TrimSpace(name) == "" {
		errors["name"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(name) > 100 {
		errors["name"] = "This field is too long (max 100 characters)"
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.plants.Insert(name)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/plant/%d", id), http.StatusSeeOther)
}

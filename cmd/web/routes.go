package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	// Dynamic routes with sessions
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/plant/create", dynamicMiddleware.ThenFunc(app.createPlantForm))
	mux.Post("/plant/create", dynamicMiddleware.ThenFunc(app.createPlant))
	mux.Get("/plant/:id", dynamicMiddleware.ThenFunc(app.showPlant))

	// Static file serving
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}

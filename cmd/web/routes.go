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

	// Plants
	mux.Get("/plant/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createPlantForm))
	mux.Post("/plant/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createPlant))
	mux.Get("/plant/:id", dynamicMiddleware.ThenFunc(app.showPlant))

	// Users
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// Static file serving
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}

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
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/scratches/create", dynamicMiddleware.ThenFunc(app.createScratchpadForm))
	mux.Post("/scratches/create", dynamicMiddleware.ThenFunc(app.createScratchpad))
	mux.Get("/scratches/:id", dynamicMiddleware.ThenFunc(app.showScratchpad))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}

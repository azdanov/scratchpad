package main

import (
	"net/http"

	"github.com/azdanov/scratchpad/ui"

	"github.com/bmizerany/pat"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, csrfCookie, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))
	mux.Get("/scratches/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createScratchpadForm))
	mux.Post("/scratches/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createScratchpad))
	mux.Get("/scratches/:id", dynamicMiddleware.ThenFunc(app.showScratchpad))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Get("/static/", fileServer)

	return standardMiddleware.Then(mux)
}

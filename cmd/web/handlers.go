package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/azdanov/scratchpad/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.scratches.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Scratches: s,
	})
}

func (app *application) showScratchpad(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.scratches.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Scratch: s,
	})
}

func (app *application) createScratchpadForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createScratchpad(w http.ResponseWriter, r *http.Request) {
	title := "Lorem ipsum"
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	expires := "7"

	id, err := app.scratches.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/scratches/%d", id), http.StatusSeeOther)
}

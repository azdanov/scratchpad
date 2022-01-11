package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	formErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		formErrors["title"] = "Title cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		formErrors["title"] = "Title is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		formErrors["content"] = "Content cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		formErrors["expires"] = "Expires cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		formErrors["expires"] = "Expires is invalid"
	}

	if len(formErrors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: formErrors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.scratches.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/scratches/%d", id), http.StatusSeeOther)
}

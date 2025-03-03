package main

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-playground/form/v4"
)

func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	component templ.Component) {

	w.WriteHeader(status)
	err := component.Render(r.Context(), w)
	if err != nil {
		app.logError(r, err)
		http.Error(w, http.StatusText(status), status)
	}

}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecodeError *form.InvalidDecoderError
		if errors.As(err, &invalidDecodeError) {
			panic(err)
		}

		return err
	}

	return nil
}

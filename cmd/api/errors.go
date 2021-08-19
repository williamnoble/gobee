package main

import "net/http"

var errorLogEntry map[string]string

func (app *application) logError(r *http.Request, err error) {
	e := errorLogEntry
	errorLogEntry["request_method"] = r.Method
	errorLogEntry["request_url"] = r.URL.String()
	app.logger.PrintError(err, e)
}


func (app *application) errorResponseTemplate(w http.ResponseWriter, r *http.Request, status int, data interface{}){
	wrapped := wrap{"error": data}
	err := write(w, status, wrapped, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

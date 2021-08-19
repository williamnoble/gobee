package main

import (
	"errors"
	"net/http"
	"sync"
)

type requestResponse int8

const (
	responseMethodNotAllowed requestResponse = iota
	responseInternalServerError
	responseResourceNotFound
	responseEditConflictResponse
	responseRateLimitExceeded
	responseInvalidAuthentication

)

var ErrWrongResponseCode = errors.New("Error in response code ")

type responseMetadata struct {
	message string
	status  int
	data    string
}

var statusList = map[requestResponse]responseMetadata{
	responseInternalServerError: {
		message: "the server encountered an error and could not process your request",
		status:  http.StatusInternalServerError,
		data:    "",
	},
	responseMethodNotAllowed: {
		message: "the method is not allowed for this resource",
		status:  http.StatusInternalServerError,
		data:    "r.Method",
	},
	responseResourceNotFound: {
		message: "the requested resource could not be found",
		status:  http.StatusInternalServerError,
		data:    "",
	},
	responseEditConflictResponse: {
		message: "unable to update the record due to an edit conflict, please try again",
		status:  http.StatusInternalServerError,
		data:    "",
	},
	responseRateLimitExceeded: {
		message: "sorry rate limit exceeed, please try again shortly",
		status:  http.StatusInternalServerError,
		data:    "",
	},
	responseInvalidAuthentication: {
		message: "sorry rate limit exceeed, please try again shortly",
		status:  http.StatusInternalServerError,
		data:    "",
	},
}

type responseMap struct {
	m  map[requestResponse]responseMetadata
	mu sync.Mutex
}

func (app *application) generateErrorResponse(w http.ResponseWriter, r *http.Request, errorCode requestResponse, err error) {

	resp := responseMap{
		m:  statusList,
		mu: sync.Mutex{},
	}


	// todo: add switch for r.Method Data
	resp.mu.Lock()

	_, ok := resp.m[errorCode]
	if !ok {
		app.logger.PrintFatal(ErrWrongResponseCode, nil)
	}

	status, msg := resp.m[errorCode].status, resp.m[errorCode].message
	// If it's an internal error or a wrong method we need to perform additional tasks
	switch errorCode {
	case responseInternalServerError:
		app.logger.PrintFatal(err, nil)
	default:
	}
	defer resp.mu.Unlock()
	app.errorResponseTemplate(w, r, status, msg)

}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponseTemplate(w, r, http.StatusInternalServerError, err)
}

func (app *application) responseMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	app.generateErrorResponse(w,r, responseMethodNotAllowed, nil)
}

func (app *application) responseNotFound(w http.ResponseWriter, r *http.Request) {
	app.generateErrorResponse(w,r, responseMethodNotAllowed, nil)
}
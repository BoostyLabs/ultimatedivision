package controllers

import (
	"net/http"
	"net/url"
)

// Redirect redirects to specific url.
func (controller *Admins) Redirect(w http.ResponseWriter, r *http.Request, urlString, method string) {
	newRequest := new(http.Request)
	*newRequest = *r
	newRequest.URL = new(url.URL)
	*newRequest.URL = *r.URL
	newRequest.Method = method

	http.Redirect(w, newRequest, urlString, http.StatusMovedPermanently)
}

package controllers

import "net/http"

// Redirect redirects to specific url.
func (controller *Users) Redirect(w http.ResponseWriter, r *http.Request, urlString, method string) {
	newRequest := r
	newRequest.URL = r.URL
	newRequest.Method = method

	http.Redirect(w, newRequest, urlString, http.StatusMovedPermanently)
}

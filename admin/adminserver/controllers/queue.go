// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/pagination"
	"ultimatedivision/queue"
)

var (
	// ErrQueue is an internal error type for queue controller.
	ErrQueue = errs.Class("queue controller error")
)

// QueueTemplates holds all queue related templates.
type QueueTemplates struct {
	List *template.Template
	Get  *template.Template
}

// Queue is a mvc controller that handles all queue related views.
type Queue struct {
	log       logger.Logger
	queue     *queue.Service
	templates QueueTemplates
}

// NewQueue is a constructor for queue controller.
func NewQueue(log logger.Logger, queue *queue.Service, templates QueueTemplates) *Queue {
	queueController := &Queue{
		log:       log,
		queue:     queue,
		templates: templates,
	}

	return queueController
}

// ListPaginated is an endpoint that will provide a web page with places.
func (controller *Queue) ListPaginated(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		placesPage  queue.Page
		err         error
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}
	placesPage, err = controller.queue.ListPaginated(ctx, cursor)
	if err != nil {
		controller.log.Error("could not list places", ErrQueue.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, placesPage)
	if err != nil {
		controller.log.Error("can not execute list places template", ErrQueue.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get is an endpoint that will provide a web page with place by id.
func (controller *Queue) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	place, err := controller.queue.Get(ctx, id)
	if err != nil {
		controller.log.Error("could not get place by id", ErrQueue.Wrap(err))
		switch {
		case queue.ErrNoPlace.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = controller.templates.Get.Execute(w, place)
	if err != nil {
		controller.log.Error("can not execute get place template", ErrQueue.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

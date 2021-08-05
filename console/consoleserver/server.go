// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/console/consoleserver/controllers"
	"ultimatedivision/internal/logger"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("console web server error")
)

// Config contains configuration for console web server.
type Config struct {
	Address string `json:"address"`
}

// Server represents console web server.
//
// architecture: Endpoint
type Server struct {
	log    logger.Logger
	config Config

	listener net.Listener
	server   http.Server
}

// NewServer is a constructor for console web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, cards *cards.Service, clubs *clubs.Service) (*Server, error) {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	cardsController := controllers.NewCards(log, cards)
	clubsController := controllers.NewClubs(log, clubs)

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	cardsRouter := apiRouter.PathPrefix("/cards").Subrouter()
	cardsRouter.HandleFunc("", cardsController.List).Methods(http.MethodGet)

	clubsRouter := apiRouter.PathPrefix("/clubs").Subrouter()
	clubsRouter.HandleFunc("{userId}", clubsController.Create).Methods(http.MethodPost)
	clubsRouter.HandleFunc("{userId}", clubsController.Get).Methods(http.MethodGet)
	clubsRouter.HandleFunc("", clubsController.UpdateSquad).Methods(http.MethodPut)

	squadsRouter := clubsRouter.Path("/squads").Subrouter()
	squadsRouter.HandleFunc("{clubId}", clubsController.Create).Methods(http.MethodPost)

	squadCardsRouter := squadsRouter.Path("/squad-cards").Subrouter()
	squadCardsRouter.HandleFunc("", clubsController.Add).Methods(http.MethodPost)
	squadCardsRouter.HandleFunc("", clubsController.UpdatePosition).Methods(http.MethodPut)
	squadCardsRouter.HandleFunc("", clubsController.Delete).Methods(http.MethodDelete)

	server.server = http.Server{
		Handler: router,
	}

	return server, nil
}

// Run starts the server that host webapp and api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return server.server.Shutdown(context.Background())
	})
	group.Go(func() error {
		defer cancel()
		err := server.server.Serve(server.listener)
		isCancelled := errs.IsFunc(err, func(err error) bool { return errors.Is(err, context.Canceled) })
		if isCancelled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return server.server.Close()
}

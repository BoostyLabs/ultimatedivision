package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
)

var (
	// ErrSecondaryMarketplace is an internal error type for secondary marketplace controller.
	ErrSecondaryMarketplace = errs.Class("secondary marketplace controller error")
)

// SecondaryMarketplace is a mvc controller that handles all secondary marketplace related views.
type SecondaryMarketplace struct {
	log         logger.Logger
	marketplace *marketplace.Service
}

// NewSecondaryMarketplace is a constructor for secondary marketplace controller.
func NewSecondaryMarketplace(log logger.Logger, marketplace *marketplace.Service) *SecondaryMarketplace {
	marketplaceController := &SecondaryMarketplace{
		log:         log,
		marketplace: marketplace,
	}

	return marketplaceController
}

// Create is an endpoint that creates lot for secondary marketplace from nft.
func (controller *SecondaryMarketplace) Create(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()
	w.Header().Set("Content-Type", "application/json")

	type request struct {
		CardID uuid.UUID `json:"cardId"`
		Price  float64   `json:"price"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrSecondaryMarketplace.Wrap(err))
		return
	}

	if req.Price <= 0 {
		controller.serveError(w, http.StatusBadRequest, ErrSecondaryMarketplace.New("price is too low"))
		return
	}

	_, err := evmsignature.EthereumFloatToWeiBig(req.Price)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrSecondaryMarketplace.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *SecondaryMarketplace) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrSecondaryMarketplace.Wrap(err))
	}
}

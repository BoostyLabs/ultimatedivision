package controllers

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"time"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/pkg/sqlsearchoperators"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
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
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var (
		startPrice big.Int
		maxPrice   big.Int
	)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemIDForm := r.FormValue("itemId")
	itemID, err := uuid.Parse(itemIDForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userIDForm := r.FormValue("userId")
	userID, err := uuid.Parse(userIDForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startPriceForm := r.FormValue("startPrice")
	if _, ok := startPrice.SetString(startPriceForm, 10); !ok {
		http.Error(w, "could not scan start price into big int", http.StatusBadRequest)
	}

	maxPriceForm := r.FormValue("maxPrice")
	if _, ok := maxPrice.SetString(maxPriceForm, 10); !ok {
		http.Error(w, "could not scan max price into big int", http.StatusBadRequest)
	}

	periodForm := r.FormValue("period")
	period, err := strconv.Atoi(periodForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createLot := marketplace.CreateLot{
		ItemID:     itemID,
		UserID:     userID,
		StartPrice: startPrice,
		MaxPrice:   maxPrice,
		Period:     marketplace.Period(period),
	}

	if err := createLot.ValidateCreateLot(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := controller.marketplace.CreateLot(ctx, createLot); err != nil {
		controller.log.Error("could not create lot", ErrSecondaryMarketplace.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetLotByID is an endpoint that returns lot by id.
func (controller *SecondaryMarketplace) GetLotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrSecondaryMarketplace.Wrap(err))
		return
	}

	lot, err := controller.marketplace.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get lot by id", ErrSecondaryMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrSecondaryMarketplace.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrSecondaryMarketplace.Wrap(err))
		}
		return
	}

	getLot := struct {
		ID           uuid.UUID          `json:"id"`
		ItemID       uuid.UUID          `json:"itemId"`
		Type         marketplace.Type   `json:"type"`
		Status       marketplace.Status `json:"status"`
		StartPrice   big.Int            `json:"startPrice"`
		MaxPrice     big.Int            `json:"maxPrice"`
		CurrentPrice big.Int            `json:"currentPrice"`
		StartTime    time.Time          `json:"startTime"`
		EndTime      time.Time          `json:"endTime"`
		Period       marketplace.Period `json:"period"`
		Card         cards.Card         `json:"card"`
	}{
		ID:           lot.ID,
		ItemID:       lot.ItemID,
		Type:         lot.Type,
		Status:       lot.Status,
		StartPrice:   lot.StartPrice,
		MaxPrice:     lot.MaxPrice,
		CurrentPrice: lot.CurrentPrice,
		StartTime:    lot.StartTime,
		EndTime:      lot.EndTime,
		Period:       lot.Period,
		Card:         lot.Card,
	}

	if err = json.NewEncoder(w).Encode(getLot); err != nil {
		controller.log.Error("failed to write json response", ErrSecondaryMarketplace.Wrap(err))
		return
	}
}

// ListLots is an endpoint that returns active lots list.
func (controller *SecondaryMarketplace) ListLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var (
		lotsPage    marketplace.Page
		err         error
		filters     cards.SliceFilters
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")
	playerName := urlQuery.Get(string(cards.FilterPlayerName))

	if limitQuery != "" {
		if limit, err = strconv.Atoi(limitQuery); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if pageQuery != "" {
		if page, err = strconv.Atoi(pageQuery); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}
	if playerName == "" {
		if err := filters.DecodingURLParameters(urlQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrSecondaryMarketplace.Wrap(err))
		}
		if len(filters) > 0 {
			lotsPage, err = controller.marketplace.ListActiveLotsWithFilters(ctx, filters, cursor)
		} else {
			lotsPage, err = controller.marketplace.ListActiveLots(ctx, cursor)
		}
	} else {
		filter := cards.Filters{
			Name:           cards.FilterPlayerName,
			Value:          playerName,
			SearchOperator: sqlsearchoperators.LIKE,
		}
		lotsPage, err = controller.marketplace.ListActiveLotsByPlayerName(ctx, filter, cursor)
	}
	if err != nil {
		controller.log.Error("could not get active lots list", ErrSecondaryMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrSecondaryMarketplace.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrSecondaryMarketplace.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(lotsPage); err != nil {
		controller.log.Error("failed to write json response", ErrSecondaryMarketplace.Wrap(err))
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

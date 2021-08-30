// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"github.com/zeebo/errs"
	"html/template"

	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

// ErrLootBoxes is an internal error type for loot boxes controller.
var ErrLootBoxes = errs.Class("lootboxes controller error")

// LootBoxesTemplates holds all lootboxes related templates.
type LootBoxesTemplates struct {
	List   *template.Template
	Create *template.Template
	Update *template.Template
}

// LootBoxes is a mvc controller that handles all lootboxes related views.
type LootBoxes struct {
	log logger.Logger

	lootboxes *lootboxes.Service

	templates LootBoxesTemplates
}

// NewLootBoxes is a constructor for loot boxes controller.
func NewLootBoxes(log logger.Logger, lootboxes *lootboxes.Service, templates LootBoxesTemplates) *LootBoxes {
	lootBoxesController := &LootBoxes{
		log:       log,
		lootboxes: lootboxes,
		templates: templates,
	}

	return lootBoxesController
}


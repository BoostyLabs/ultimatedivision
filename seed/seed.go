// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seed

import (
	"context"
	"time"

	"github.com/google/uuid"

	"ultimatedivision"
	"ultimatedivision/admin/admins"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/seasons"
	"ultimatedivision/users"
)



// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Admins provides access to admins db.
	Admins() admins.DB

	// Users provides access to users db.
	Users() users.DB

	// Cards provides access to cards db.
	Cards() cards.DB

	// Clubs provides access to clubs db.
	Clubs() clubs.DB

	// Matches provides access to matches db.
	Matches() matches.DB

	// Divisions provides access to divisions db.
	Divisions() divisions.DB

	// Seasons provides access to seasons db.
	Seasons() seasons.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) error
}

// Seed fills database table with data.
type Seed struct {
	Config   ultimatedivision.Config
	Database DB

	Admins struct {
		Service *admins.Service
	}

	Users struct {
		Service *users.Service
	}

	Clubs struct {
		Service *clubs.Service
	}

	Cards struct {
		Service *cards.Service
	}

	Matches struct {
		Service *matches.Service
	}

	Divisions struct {
		Service *divisions.Service
	}

	Seasons struct {
		Service *seasons.Service
	}
}

// New is constructor for Seed.
func New(database DB, config ultimatedivision.Config) *Seed {
	seed := &Seed{
		Database: database,
		Config:   config,
	}

	seed.Admins.Service = admins.NewService(
		seed.Database.Admins(),
	)

	seed.Users.Service = users.NewService(
		seed.Database.Users(),
	)

	seed.Cards.Service = cards.NewService(
		seed.Database.Cards(),
		seed.Config.Cards.Config,
	)

	seed.Divisions.Service = divisions.NewService(
		seed.Database.Divisions(),
		seed.Config.Divisions.Config,
	)

	seed.Clubs.Service = clubs.NewService(
		seed.Database.Clubs(),
		seed.Users.Service,
		seed.Cards.Service,
		seed.Database.Divisions(),
	)

	seed.Matches.Service = matches.NewService(
		seed.Database.Matches(),
		seed.Config.Matches.Config,
		seed.Clubs.Service,
	)

	seed.Seasons.Service = seasons.NewService(
		seed.Database.Seasons(),
		seed.Config.Seasons.Config,
		seed.Divisions.Service,
		seed.Matches.Service,
		seed.Clubs.Service,
	)

	return seed
}

// Seed fills user, admin, club, squad and squad cards table with test value.s
func(seed *Seed) Seed(ctx context.Context) error {
	// creating users
	allUsers := createUsers()
	for _, user := range allUsers {
		err := seed.Users.Service.Create(ctx, user.Email, string(user.PasswordHash), user.NickName, user.FirstName, user.LastName)
		if err != nil {
			return err
		}
	}

	// creating admin
	admin := admins.Admin{
		Email : "test@test.com",
		PasswordHash: []byte("Qwerty123-"),
	}
	err := seed.Admins.Service.Create(ctx, admin.Email, admin.PasswordHash)
	if err != nil {
		return err
	}

	// creating divisions.
	allNames := []int{1,2,3,4,5,6,7,8,9,10}
	for _, name := range allNames {
		err := seed.Divisions.Service.Create(ctx, name)
		if err != nil {
			return err
		}
	}

	// creating clubs
	allUsers, err = seed.Users.Service.List(ctx)
	if err != nil {
		return err
	}

	for _, user := range allUsers {
		_, err = seed.Clubs.Service.Create(ctx, user.ID)
		if err != nil {
			return err
		}
	}

	// creating squads for clubs.
	allClubs, err := seed.Clubs.Service.List(ctx)
	if err != nil {
		return err
	}

	for _, club := range allClubs {
		_, err := seed.Clubs.Service.CreateSquad(ctx, club.ID)
		if err != nil {
			return err
		}
	}

	// creating cards for squad.
	allClubs, err = seed.Clubs.Service.List(ctx)
	if err != nil {
		return err
	}

	var squadCards []clubs.SquadCard

	for _, club := range allClubs {
		squad, err := seed.Clubs.Service.GetSquadByClubID(ctx, club.ID)
		if err != nil {
			return err
		}
		probabilities := []int{seed.Config.LootBoxes.Config.RegularBoxConfig.Wood, seed.Config.LootBoxes.Config.RegularBoxConfig.Silver, seed.Config.LootBoxes.Config.RegularBoxConfig.Gold, seed.Config.LootBoxes.Config.RegularBoxConfig.Diamond}
		for i := 0; i < 11 ; i++ {
			card, err := seed.Cards.Service.Create(ctx, club.OwnerID, probabilities)
			if err != nil {
				return err
			}
			squadCard := clubs.SquadCard{
				SquadID:  squad.ID,
				CardID:   card.ID,
				Position: clubs.Position(i),
			}

			squadCards = append(squadCards, squadCard)
		}
	}

	for _, card := range squadCards {
		err = seed.Clubs.Service.AddSquadCard(ctx, card.SquadID, card)
		if err != nil {
			return err
		}
	}

	return nil
}

// createUsers creates slice of users.
func createUsers() []users.User {
	testUser1 := users.User{
		ID:           uuid.New(),
		Email:        "testUser1@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin1",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser2 := users.User{
		ID:           uuid.New(),
		Email:        "testUser2@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin2",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser3 := users.User{
		ID:           uuid.New(),
		Email:        "testUser3@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin3",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser4 := users.User{
		ID:           uuid.New(),
		Email:        "testUser4@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin4",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser5 := users.User{
		ID:           uuid.New(),
		Email:        "testUser5@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin5",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	testUser6 := users.User{
		ID:           uuid.New(),
		Email:        "testUser6@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin6",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	return []users.User{testUser1, testUser2, testUser3, testUser4, testUser5, testUser6}
}

// Match initiates matches between users.
func(seed *Seed) Match(ctx context.Context) error {
	type player struct {
		userID   uuid.UUID
		squadID  uuid.UUID
		seasonID int
	}

	var players []player

	allClubs, err := seed.Clubs.Service.List(ctx)
	if err != nil {
		return err
	}

	for _, club := range allClubs {
		squad, err := seed.Clubs.Service.Get(ctx, club.ID)
		if err != nil {
			return err
		}
		season, err := seed.Seasons.Service.GetSeasonByDivisionID(ctx, club.DivisionID)
		if err != nil {
			return err
		}
		player := player{
			userID:   club.OwnerID,
			squadID:  squad.ID,
			seasonID: season.ID,
		}

		players = append(players, player)
	}

	index := 1

	for _, player1 := range players {
		for _, player2 := range players[:len(players)-index] {
			_, err := seed.Matches.Service.Create(ctx, player1.squadID, player2.squadID, player1.userID, player2.userID, player1.seasonID)
			if err != nil {
				return err
			}
		}
		index++
	}

	return nil
}

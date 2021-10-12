// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoClub indicated that club does not exist.
var ErrNoClub = errs.Class("club does not exist")

// ErrNoSquad indicated that squad does not exist.
var ErrNoSquad = errs.Class("squad does not exist")

// DB is exposing access to clubs db.
//
// architecture: DB
type DB interface {
	// Create creates club in the database.
	Create(ctx context.Context, club Club) (uuid.UUID, error)
	// CreateSquad creates squad for clubs in the database.
	CreateSquad(ctx context.Context, squad Squad) (uuid.UUID, error)
	// GetByUserID returns club owned by the user.
	GetByUserID(ctx context.Context, userID uuid.UUID) (Club, error)
	// GetSquad returns squad.
	GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, error)
	// GetFormation returns formation of the squad.
	GetFormation(ctx context.Context, squadID uuid.UUID) (Formation, error)
	// GetCaptainID returns id of captain.
	GetCaptainID(ctx context.Context, squadID uuid.UUID) (uuid.UUID, error)
	// ListSquadCards returns all cards from squad.
	ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]SquadCard, error)
	// AddSquadCard adds new card to the squad.
	AddSquadCard(ctx context.Context, squadCards SquadCard) error
	// DeleteSquadCard deletes card from squad.
	DeleteSquadCard(ctx context.Context, squadID, cardID uuid.UUID) error
	// UpdateTacticFormationCaptain updates tactic, formation and capitan in the squad.
	UpdateTacticFormationCaptain(ctx context.Context, squad Squad) error
	// UpdatePositions updates positions of cards in the squad.
	UpdatePositions(ctx context.Context, squadCards []SquadCard) error
	// UpdateFormation updates formation in the squad.
	UpdateFormation(ctx context.Context, newFormation Formation, squadID uuid.UUID) error
}

// Club defines club entity.
type Club struct {
	ID        uuid.UUID `json:"id"`
	OwnerID   uuid.UUID `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Squad describes squads of clubs.
type Squad struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"-"`
	ClubID    uuid.UUID `json:"clubId"`
	Formation Formation `json:"formation"`
	Tactic    Tactic    `json:"tactic"`
	CaptainID uuid.UUID `json:"captainId"`
}

// SquadCard defines all cards from squad.
type SquadCard struct {
	SquadID  uuid.UUID `json:"squadId"`
	CardID   uuid.UUID `json:"cardId"`
	Position Position  `json:"position"`
}

// Formation defines a list of possible formations.
type Formation int

const (
	// FourFourTwo defines 4-4-2 scheme.
	FourFourTwo Formation = 1
	// FourTwoFour defines 4-2-4 scheme.
	FourTwoFour Formation = 2
	// FourTwoTwoTwo defines 4-2-2-2 scheme.
	FourTwoTwoTwo Formation = 3
	// FourThreeOneTwo defines 4-3-1-2 scheme.
	FourThreeOneTwo Formation = 4
	// FourThreeThree defines 4-3-3 scheme.
	FourThreeThree Formation = 5
	// FourTwoThreeOne defines 4-2-3-1 scheme.
	FourTwoThreeOne Formation = 6
	// FourThreeTwoOne defines 4-3-2-1 scheme.
	FourThreeTwoOne Formation = 7
	// FourOneThreeTwo defines 4-1-3-2 scheme.
	FourOneThreeTwo Formation = 8
	// FiveThreeTwo defines 5-3-2 scheme.
	FiveThreeTwo Formation = 9
	// ThreeFiveTwo defines 4-5-2 scheme.
	ThreeFiveTwo Formation = 10
)

// IsValid check that formation ID is valid.
func (f Formation) IsValid() bool {
	switch f {
	case FourFourTwo,
		FourTwoFour,
		FourTwoTwoTwo,
		FourThreeOneTwo,
		FourThreeThree,
		FourTwoThreeOne,
		FourThreeTwoOne,
		FourOneThreeTwo,
		FiveThreeTwo,
		ThreeFiveTwo:
		return true
	default:
		return false
	}
}

// Tactic defines a list of possible tactics.
type Tactic int

const (
	// Attack defines attacking style.
	Attack Tactic = 1
	// Defence defines defensive style.
	Defence Tactic = 2
	// Balanced balance between attack and defense.
	Balanced Tactic = 3
)

// Position defines a list of possible positions.
type Position int

const (
	// GK defines goalkeeper.
	GK Position = 1
	// LB defines left defenders.
	LB Position = 2
	// LCD defines left central defenders.
	LCD Position = 3
	// CCD defines center central defenders.
	CCD Position = 4
	// RCD defines right central defenders.
	RCD Position = 5
	// RB defines right defenders.
	RB Position = 6
	// LCDM defines left central defensive midfielder.
	LCDM Position = 7
	// CCDM defines center central defensive midfielder.
	CCDM Position = 8
	// RCDM defines right central defensive midfielder.
	RCDM Position = 9
	// LCM defines left central midfielder.
	LCM Position = 10
	// CCM defines central central midfielder.
	CCM Position = 11
	// RCM defines right central midfielder.
	RCM Position = 12
	// LM defines left midfielder.
	LM Position = 13
	// RM defines right midfielder.
	RM Position = 14
	// LCAM defines left central attacking midfielder.
	LCAM Position = 15
	// CCAM defines center central attacking midfielder.
	CCAM Position = 16
	// RCAM defines right central attacking midfielder.
	RCAM Position = 17
	// LWB defines left attacking defenders.
	LWB Position = 18
	// RWB defines right attacking defenders.
	RWB Position = 19
	// RW defines right forward.
	RW Position = 20
	// LW defines left forward.
	LW Position = 21
	// LST defines left central forward.
	LST Position = 22
	// RST defines right central forward.
	RST Position = 23
	// CST defines center central forward.
	CST Position = 24
)

// FormationToPosition defines positions that are present in the formation.
var FormationToPosition = map[Formation][]Position{
	FourFourTwo:     {GK, LB, LCD, RCD, RB, LM, LCM, RCM, RM, LST, RST},
	FourTwoFour:     {GK, LB, LCD, RCD, RB, LCM, RCM, LW, LST, RST, RW},
	FourTwoTwoTwo:   {GK, LB, LCD, RCD, RB, LCAM, LCDM, RCDM, RCAM, LST, RST},
	FourThreeOneTwo: {GK, LB, LCD, RCD, RB, LCM, CCM, CCAM, RCM, LST, RST},
	FourThreeThree:  {GK, LB, LCD, RCD, RB, LCM, CCM, RCM, LW, CST, RW},
	FourTwoThreeOne: {GK, LB, LCD, RCD, RB, LCDM, LCAM, CCAM, RCAM, RCDM, CST},
	FourThreeTwoOne: {GK, LB, LCD, RCD, RB, LCM, CCM, RCM, LW, CST, RW},
	FourOneThreeTwo: {GK, LB, LCD, RCD, RB, LM, CCM, CCDM, RM, LST, RST},
	FiveThreeTwo:    {GK, LWB, LCD, CCD, RCD, RWB, LCM, CCM, RCM, LST, RST},
	ThreeFiveTwo:    {GK, LCD, CCD, RCD, LM, LCDM, CCAM, RCDM, RM, LST, RST},
}

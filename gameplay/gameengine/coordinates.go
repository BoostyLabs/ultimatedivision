// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
)

// Coordinate defines dot in coordinate plane.
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Compare compares two coordinates.
func (coordinate Coordinate) Compare(coordinateToCompare Coordinate) bool {
	return coordinate.X == coordinateToCompare.X && coordinate.Y == coordinateToCompare.Y
}

// CoordinatesConfig contains all config values related to field coordinates.
type CoordinatesConfig struct {
	SizeOfFieldByOX int `json:"sizeOfFieldByOX"`
	SizeOfFieldByOY int `json:"sizeOfFieldByOY"`

	Positions struct {
		GK struct {
			Coordinate
		} `json:"gk"`
		LB struct {
			Coordinate
		} `json:"lb"`
		LCB struct {
			Coordinate
		} `json:"lcb"`
		CCB struct {
			Coordinate
		} `json:"ccb"`
		RCB struct {
			Coordinate
		} `json:"rcb"`
		RB struct {
			Coordinate
		} `json:"rb"`
		LCDM struct {
			Coordinate
		} `json:"lcdm"`
		CCDM struct {
			Coordinate
		} `json:"ccdm"`
		RCDM struct {
			Coordinate
		} `json:"rcdm"`
		LCM struct {
			Coordinate
		} `json:"lcm"`
		CCM struct {
			Coordinate
		} `json:"ccm"`
		RCM struct {
			Coordinate
		} `json:"rcm"`
		LM struct {
			Coordinate
		} `json:"lm"`
		RM struct {
			Coordinate
		} `json:"rm"`
		LCAM struct {
			Coordinate
		} `json:"lcam"`
		CCAM struct {
			Coordinate
		} `json:"ccam"`
		RCAM struct {
			Coordinate
		} `json:"rcam"`
		LWB struct {
			Coordinate
		} `json:"lwb"`
		RWB struct {
			Coordinate
		} `json:"rwb"`
		LW struct {
			Coordinate
		} `json:"lw"`
		RW struct {
			Coordinate
		} `json:"rw"`
		LST struct {
			Coordinate
		} `json:"lst"`
		CST struct {
			Coordinate
		} `json:"cst"`
		RST struct {
			Coordinate
		} `json:"rst"`
	} `json:"positions"`
}

// CardWithCoordinate defines card with coordinate in the field.
type CardWithCoordinate struct {
	cards.Card `json:"card"`
	Coordinate Coordinate `json:"coordinate"`
}

// ConvertPositionToCoordinate converts positions to coordinates.
func ConvertPositionToCoordinate(config CoordinatesConfig, squadCards []clubs.SquadCard) []CardWithCoordinate {
	var PositionToCoordinates = map[clubs.Position]Coordinate{
		clubs.GK:   config.Positions.GK.Coordinate,
		clubs.LB:   config.Positions.LB.Coordinate,
		clubs.LCD:  config.Positions.LCB.Coordinate,
		clubs.CCD:  config.Positions.CCB.Coordinate,
		clubs.RCD:  config.Positions.RCB.Coordinate,
		clubs.RB:   config.Positions.RB.Coordinate,
		clubs.LCDM: config.Positions.LCDM.Coordinate,
		clubs.RCDM: config.Positions.RCDM.Coordinate,
		clubs.CCDM: config.Positions.CCDM.Coordinate,
		clubs.CCM:  config.Positions.CCM.Coordinate,
		clubs.RCM:  config.Positions.RCM.Coordinate,
		clubs.LCM:  config.Positions.LCM.Coordinate,
		clubs.LM:   config.Positions.LM.Coordinate,
		clubs.RM:   config.Positions.RM.Coordinate,
		clubs.CCAM: config.Positions.CCAM.Coordinate,
		clubs.RCAM: config.Positions.RCAM.Coordinate,
		clubs.LCAM: config.Positions.LCAM.Coordinate,
		clubs.LWB:  config.Positions.LWB.Coordinate,
		clubs.RWB:  config.Positions.RWB.Coordinate,
		clubs.LW:   config.Positions.LW.Coordinate,
		clubs.RW:   config.Positions.RW.Coordinate,
		clubs.RST:  config.Positions.RST.Coordinate,
		clubs.CST:  config.Positions.CST.Coordinate,
		clubs.LST:  config.Positions.LST.Coordinate,
	}

	var cardsWithCoordinate []CardWithCoordinate

	for _, squadCard := range squadCards {
		var cardWithCoordinate CardWithCoordinate

		cardWithCoordinate.Coordinate = PositionToCoordinates[squadCard.Position]
		cardWithCoordinate.Card = squadCard.Card
		cardsWithCoordinate = append(cardsWithCoordinate, cardWithCoordinate)
	}

	return cardsWithCoordinate
}

// ReflectCoordinates reflects coordinates relative to the center of the field.
func ReflectCoordinates(cardsWithCoordinate []CardWithCoordinate, sizeOfFieldByOX, sizeOfFieldByOY int) []CardWithCoordinate {
	var cardsWithNewCoordinates []CardWithCoordinate

	for _, cardWithCoordinate := range cardsWithCoordinate {
		var cardWithNewCoordinate CardWithCoordinate
		cardWithNewCoordinate.Card = cardWithCoordinate.Card

		if cardWithCoordinate.Coordinate.X > sizeOfFieldByOX/2 {
			cardWithNewCoordinate.Coordinate.X -= sizeOfFieldByOY / 2
			continue
		}
		cardWithNewCoordinate.Coordinate.X += sizeOfFieldByOX - cardWithCoordinate.Coordinate.X

		cardsWithNewCoordinates = append(cardsWithNewCoordinates, cardWithNewCoordinate)
	}
	return cardsWithNewCoordinates
}

// GetCenterOfField returns coordinate of the center of field.
func GetCenterOfField(sizeOfFieldByOX, sizeOfFieldByOY int) Coordinate {
	return Coordinate{
		X: sizeOfFieldByOX / 2,
		Y: sizeOfFieldByOY / 2,
	}
}

// GenerateCellsInRange returns all cells in specific radius from start coordinate.
func generateCellsInRange(startCoordinate Coordinate, numOfCells, sizeOfFieldByOX, sizeOfFieldByOY int, opponentCardsWithPositions []Coordinate) []Coordinate {
	var cells []Coordinate

	minCoordinateOX := startCoordinate.X - numOfCells
	minCoordinateOY := startCoordinate.Y - numOfCells
	maxCoordinateOX := startCoordinate.X + numOfCells
	maxCoordinateOY := startCoordinate.Y + numOfCells

	switch {
	case minCoordinateOX > sizeOfFieldByOX:
		minCoordinateOX = sizeOfFieldByOX
	case minCoordinateOX < 0:
		minCoordinateOX = 0
	case minCoordinateOY > sizeOfFieldByOY:
		minCoordinateOY = sizeOfFieldByOY
	case minCoordinateOY < 0:
		minCoordinateOY = 0
	case maxCoordinateOX > sizeOfFieldByOX:
		maxCoordinateOX = sizeOfFieldByOX
	case maxCoordinateOX < 0:
		maxCoordinateOX = 0
	case maxCoordinateOY > sizeOfFieldByOY:
		maxCoordinateOY = sizeOfFieldByOY
	case maxCoordinateOY < 0:
		maxCoordinateOY = 0
	}

	for i := minCoordinateOX; i <= maxCoordinateOX; i++ {
		for j := minCoordinateOY; j <= maxCoordinateOY; j++ {
			if i == sizeOfFieldByOX || j == sizeOfFieldByOY || i < 0 || j < 0 {
				continue
			}

			cells = append(cells, Coordinate{
				X: i,
				Y: j,
			})
		}
	}

	cells = deleteElementWhichExistsInArray(cells, opponentCardsWithPositions)

	return cells
}

// DeleteElementWhichExistsInArray deletes element from allies slice of cards with coordinates
// where coordinate coordinates are identical.
func deleteElementWhichExistsInArray(alliesCardsWithPositions, opponentCardsWithPositions []Coordinate) []Coordinate {
	for i := 0; i < len(alliesCardsWithPositions); i++ {
		for j := 0; j < len(opponentCardsWithPositions); j++ {
			if alliesCardsWithPositions[i].Compare(opponentCardsWithPositions[j]) {
				alliesCardsWithPositions = removeElementFromSlice(alliesCardsWithPositions, i)
			}
		}
	}

	return alliesCardsWithPositions
}

// removeElementFromSlice removes element from slice with certain index.
func removeElementFromSlice(cardsWithPosition []Coordinate, index int) []Coordinate {
	return append(cardsWithPosition[:index], cardsWithPosition[index+1:]...)
}

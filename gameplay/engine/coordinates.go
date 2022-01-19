// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package engine

// Coordinate defines coordinate of card/ball in the field.
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Compare compares two coordinates.
func (c Coordinate) Compare(comparable Coordinate) bool {
	return c.X == comparable.X && c.Y == comparable.Y
}

type Config struct {
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

	MoveAction struct {
		NumOfCellsForFirstRange  int `json:"numOfCellsForFirstRange"`
		NumOfCellsForSecondRange int `json:"numOfCellsForSecondRange"`
		NumOfCellsForThirdRange  int `json:"numOfCellsForThirdRange"`
		NumOfCellsForFourthRange int `json:"numOfCellsForFourthRange"`

		FirstRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"firstRange"`
		SecondRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"secondRange"`
		ThirdRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"thirdRange"`
		FourthRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"fourthRange"`
	} `json:"moveAction"`
}


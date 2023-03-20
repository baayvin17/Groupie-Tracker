package _var

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

var (
	MinMaxBool     bool
	MinMaxDatelist []string
	MinMaxTimelist []string
	SelectLoclist  []string
	TimeBool       bool
	Layout         = "02-01-2006"
	SliderBool     bool

	Inputsaisie *widget.Entry                     // contenu de la bar de recherche
	Resultat    []string                          // contenue du Resultat de la recherche
	EntryResult = widget.NewSelect(Resultat, nil) // Resultat de la recherche

	UserInputInt int
	UserMin      int
	UserMinDate  time.Time
	UserMax      int
	UserMaxDate  time.Time

	BoolBand    bool // pr les check-list
	BoolMembers bool
	BoolLoc     bool
	BoolConc    bool
	BoolCrea    bool
	BoolAlbm    bool

	UserNumMembers int
	UserLoc        string
	SelectLoc      bool
)

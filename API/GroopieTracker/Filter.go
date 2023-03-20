package GroopieTracker

import (
	"fmt"
	_var "main/res/var"

	"time"
)

// TODO test des deux types de filtres ensemble sur fyne

func FilterApi() {
	Checklist()
	MinMax()
	Slider()
}

func Checklist() {
	_var.Resultat = []string{}

	if _var.BoolBand { // letters
		fmt.Println(_var.Inputsaisie.Text)
		for _, s := range SearchBand(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
	if _var.BoolMembers == true { // letters
		fmt.Println("Boolmembers : ", _var.BoolMembers)
		for _, s := range SearchMember(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
	if _var.BoolLoc == true { // letters
		fmt.Println("BoolLoc : ", _var.BoolLoc)
		for _, s := range SearchLocation(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
	if _var.BoolConc == true { // numbers
		fmt.Println("BoolConc : ", _var.BoolConc)
		for _, s := range SearchConcDate(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
	if _var.BoolCrea == true { // numbers
		fmt.Println("BoolCrea : ", _var.BoolCrea)
		for _, s := range SearchCreaDate(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
	if _var.BoolAlbm == true { // numbers
		fmt.Println("BoolAlbm", _var.BoolAlbm)
		for _, s := range SearchAlbumDate(_var.Inputsaisie.Text) {
			_var.Resultat = append(_var.Resultat, s)
		}
	}
}
func MinMax() {
	// debug false = creaDate, true = albDate + concDate
	if _var.TimeBool == false && _var.MinMaxBool == true {
		_var.Resultat = append(_var.Resultat, SearchCreaDate(_var.Inputsaisie.Text)...)
	}
	if _var.TimeBool == true && _var.MinMaxBool == false {
		_var.UserMinDate, _ = time.Parse(_var.Layout, _var.Inputsaisie.Text)
		_var.UserMaxDate, _ = time.Parse(_var.Layout, _var.Inputsaisie.Text)

		_var.Resultat = append(_var.Resultat, SearchCreaDate(_var.Inputsaisie.Text)...)
		_var.Resultat = append(_var.Resultat, SearchConcDate(_var.Inputsaisie.Text)...)
		_var.Resultat = append(_var.Resultat, SearchAlbumDate(_var.Inputsaisie.Text)...)
	}

}
func Slider() {
	if _var.SliderBool == true {

	}
}

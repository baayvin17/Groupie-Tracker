package Search_Bar

import (
	"main/API/GroopieTracker"
	"main/res"
	_var "main/res/var"

	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SearchBar() *widget.Entry {
	Entry := widget.NewEntry()
	Entry.SetPlaceHolder(res.L.SearchBar)
	_var.Inputsaisie = Entry
	return Entry
}

func SearchButton() *widget.Button {
	EntryBtn := widget.NewButton(res.L.BUTTONS["Search"], func() {
		GroopieTracker.SearchApi()
		GroopieTracker.MinMax()
		SearchResult()
	})
	return EntryBtn
}

func SearchResult() *widget.Select {
	_var.EntryResult.Options = _var.Resultat
	_var.EntryResult.Selected = ""
	_var.EntryResult.OnChanged = func(s string) {
		GroopieTracker.SelectFunc(s)
	}
	if len(_var.Resultat) == 0 {
		_var.EntryResult.PlaceHolder = res.L.SearchNoResultPlaceholder
	} else {
		_var.EntryResult.PlaceHolder = _var.Resultat[0] + " > " + strconv.Itoa(len(_var.Resultat)) + res.L.SearchResultPlaceholder
	}
	_var.EntryResult.Refresh()
	return _var.EntryResult
}
func FiltreCheckBox1() *fyne.Container {

	checkband := widget.NewCheck(res.L.FilterBand, func(b bool) {
		_var.BoolBand = b
	})
	checkband.Checked = true
	checkband.Refresh()
	_var.BoolBand = true

	checkmember := widget.NewCheck(res.L.FilterMember, func(m bool) {
		_var.BoolMembers = m
	})
	checkmember.Checked = true
	checkmember.Refresh()
	_var.BoolMembers = true

	checkloc := widget.NewCheck(res.L.FilterLocation, func(l bool) {
		_var.BoolLoc = l
	})
	checkloc.Checked = true
	checkloc.Refresh()
	_var.BoolLoc = true

	checkconc := widget.NewCheck(res.L.FilterConcertDate, func(c bool) {
		_var.BoolConc = c
	})
	checkconc.Checked = true
	checkband.Refresh()
	_var.BoolConc = true

	checkcrea := widget.NewCheck(res.L.FilterCreationDate, func(r bool) {
		_var.BoolCrea = r
	})
	checkcrea.Checked = true
	checkcrea.Refresh()
	_var.BoolCrea = true

	checkalbm := widget.NewCheck(res.L.FilterAlbumDate, func(a bool) {
		_var.BoolAlbm = a
	})
	checkalbm.Checked = true
	checkalbm.Refresh()
	_var.BoolAlbm = true

	HBOX := container.NewHBox(checkband, checkmember, checkloc, checkconc, checkcrea, checkalbm)

	return HBOX
}

func FiltreCheckBox2() *fyne.Container {
	radiolabel := widget.NewLabel(res.L.FilterNbMember)
	radioMembers := widget.NewRadioGroup([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, func(value string) {
		_var.UserNumMembers, _ = strconv.Atoi(value)
	})
	radioMembers.Horizontal = true

	HBOXC4 := container.NewHBox(radiolabel, radioMembers)

	return HBOXC4
}

func FiltreMinMax1() *fyne.Container {
	GroopieTracker.SearchCreaDate("")
	minlabel := widget.NewLabel(res.L.FilterMinCreationDate)
	minCreaDate := widget.NewSelect(_var.MinMaxDatelist, func(value string) {
		_var.MinMaxBool = true
		_var.TimeBool = false
		_var.UserMin, _ = strconv.Atoi(value)
	})
	maxlabel := widget.NewLabel(res.L.FilterMaxCreationDate)
	maxCreaDate := widget.NewSelect(_var.MinMaxDatelist, func(value string) {
		_var.MinMaxBool = true
		_var.TimeBool = false
		_var.UserMax, _ = strconv.Atoi(value)
	})
	HBOX2 := container.NewHBox(minlabel, minCreaDate, maxlabel, maxCreaDate)
	return HBOX2
}

func FiltreMinMax2() *fyne.Container {
	GroopieTracker.SearchConcDate("")
	GroopieTracker.SearchAlbumDate("")

	mindatelabel := widget.NewLabel(res.L.FilterMinConcertDate)
	minConcDate := widget.NewSelect(_var.MinMaxTimelist, func(value string) {
		_var.MinMaxBool = false
		_var.TimeBool = true
		_var.UserMinDate, _ = time.Parse(_var.Layout, value)
	})
	maxdatelabel := widget.NewLabel(res.L.FilterMaxConcertDate)
	maxConcDate := widget.NewSelect(_var.MinMaxTimelist, func(value string) {
		_var.MinMaxBool = false
		_var.TimeBool = true
		_var.UserMaxDate, _ = time.Parse(_var.Layout, value)
	})
	HBOX3 := container.NewHBox(mindatelabel, minConcDate, maxdatelabel, maxConcDate)
	return HBOX3
}

func FiltreMinMaxSlider() *fyne.Container {
	sliderlabel := widget.NewLabel(res.L.FilterNbMember)
	SliderMember := widget.NewSlider(0, 8)
	if SliderMember.Step > 0 {
		_var.UserNumMembers = int(SliderMember.Value)
		_var.SliderBool = true
	}
	HBOX3 := container.NewHBox(sliderlabel, SliderMember)
	return HBOX3
}

func FiltreLoc() *fyne.Container {
	selectlabel := widget.NewLabel(res.L.FilterLocation)
	selectLoc := widget.NewSelect(_var.SelectLoclist, func(value string) {
		_var.UserLoc = value
		_var.SelectLoc = true
	})
	HBOXSelect := container.NewHBox(selectlabel, selectLoc)
	return HBOXSelect
}

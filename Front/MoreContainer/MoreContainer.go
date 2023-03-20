package MoreContainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/Front/MoreContainer/MemberContainer"
	RES "main/res"
	data "main/res/var"
)

func More(artist int) func() {
	return func() {
		MoreWindow(artist)
	}
}

func MoreContainer(artiste int) *fyne.Container {

	// New container
	Container := container.NewVBox()          // Main Container
	ImageDescContainer := container.NewHBox() // Image and Desc Container
	MembersContainer := container.NewAdaptiveGrid(1)

	// Load image
	res, _ := fyne.LoadResourceFromURLString(data.AData[artiste].Image)
	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillOriginal
	img.Resize(fyne.NewSize(200, 200))

	// Load members
	members := ""
	for mem := range data.AData[artiste].Members {
		widget.NewLabel(data.AData[artiste].Members[mem])
	}

	// Load card
	Card := widget.NewCard(data.AData[artiste].Name, members, img)
	ImageDescContainer.Add(Card)

	// Desc Container
	DescContainer := container.NewVBox()

	// First Album
	DescContainer.Add(widget.NewLabel(RES.L.ArtistFirstAlbum))
	DescContainer.Add(widget.NewLabel(data.AData[artiste].FirstAlbum))

	//------------------------------------------------------------------------------------------------------------------
	DescContainer.Add(layout.NewSpacer())

	// Concerts
	DescContainer.Add(widget.NewLabel(RES.L.ArtistConcerts))
	DateLocation := container.NewHBox()
	for location, date := range data.RData.Index[artiste].DateLocation {
		h := container.NewVBox()
		LocationSelector := widget.NewSelect(date, nil)
		LocationSelector.SetSelectedIndex(-1)
		LocationSelector.PlaceHolder = date[0] + "..."

		h.Add(widget.NewLabel(location + " : "))
		h.Add(LocationSelector)
		DateLocation.Add(h)

	}
	DescContainer.Add(DateLocation)

	//------------------------------------------------------------------------------------------------------------------
	DescContainer.Add(layout.NewSpacer())

	ImageDescContainer.Add(DescContainer)

	// GeoLocation
	/*res, _ = fyne.LoadResourceFromURLString("https://lululataupe.com/images/decouverte/geographie/carte-du-monde-avec-capitales-pays.png")
	GeoLocation := canvas.NewImageFromResource(res)
	//GeoLocation.FillMode = canvas.ImageFillOriginal
	GeoLocation.Resize(fyne.NewSize(200, 200))
	Members_Container.Add(GeoLocation)

	*/

	// list of members
	MembersContainer.Add(MemberContainer.MemberList(artiste))

	Container.Add(ImageDescContainer)
	Container.Add(MembersContainer)
	return Container
}

func MoreWindow(artist int) {
	// New window
	w := fyne.CurrentApp().NewWindow(data.AData[artist].Name)
	w.Resize(fyne.NewSize(800, 600))

	//set the icon
	w.SetIcon(RES.AppIcon)

	w.SetContent(
		container.NewScroll(MoreContainer(artist)),
	)

	w.Show()
}

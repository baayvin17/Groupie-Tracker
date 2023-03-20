package MemberContainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/res"
	data "main/res/var"
)

func MemberList(artiste int) *container.Scroll {

	Container := container.NewAdaptiveGrid(len(data.AData[artiste].Members))
	for mem := range data.AData[artiste].Members {
		//load image
		resource, _ := fyne.LoadResourceFromURLString(data.AData[artiste].Image)
		img := canvas.NewImageFromResource(resource)
		img.FillMode = canvas.ImageFillOriginal

		card := widget.NewCard(data.AData[artiste].Members[mem], "", img)
		cont := container.New(layout.NewMaxLayout(), widget.NewButton("", ShowMember(artiste, mem)), card)
		Container.Add(cont)

	}
	scroll := container.NewScroll(Container)
	scroll.Direction = container.ScrollHorizontalOnly
	return scroll
}

func ShowMember(art, mem int) func() {
	return func() {
		//fix link
		OpenWiki(art, mem)
	}
}

func OpenWiki(art, mem int) {
	url := res.UrlWikipedia
	for _, char := range data.AData[art].Members[mem] {
		if char == ' ' {
			url += "_"
		} else {
			url += string(char)
		}
	}

	//open link
	res.OpenBrowser(url)
}

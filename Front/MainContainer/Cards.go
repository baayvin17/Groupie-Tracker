package MainContainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/Front/MoreContainer"
	RES "main/res"
	_var "main/res/var"
)

func Cards() *fyne.Container {
	Container := container.NewAdaptiveGrid(5)

	for artiste := range _var.AData {
		res, _ := fyne.LoadResourceFromURLString(_var.AData[artiste].Image)
		img := canvas.NewImageFromResource(res)
		img.FillMode = canvas.ImageFillOriginal
		img.Resize(fyne.NewSize(200, 200))

		members := ""
		for mem := range _var.AData[artiste].Members {
			if len(members) < 25 {
				members += _var.AData[artiste].Members[mem]
				if mem < len(_var.AData[artiste].Members)-1 {
					members += ", "
				}

			} else {
				members += "..."
				break
			}
		}
		card := widget.NewCard(_var.AData[artiste].Name, members, img)
		MainCard := container.New(layout.NewMaxLayout(), widget.NewButton(RES.L.BUTTONS["more"], MoreContainer.More(artiste)), card)
		Container.Add(MainCard)

	}
	return Container
}

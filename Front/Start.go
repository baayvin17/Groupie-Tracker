package Start

import (
	"main/Front/MainContainer"
	"main/Front/menu"
	"main/res"
	"main/res/save"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func Start() {
	// new app
	A := app.New()

	//Init the preferences
	save.InitPref()

	// New title and window
	W := A.NewWindow(res.L.AppName + " " + res.L.AppVersion)
	W.Resize(fyne.NewSize(1920, 1080))
	W.SetMaster()

	//set the icon
	W.SetIcon(res.AppIcon)

	W.SetMainMenu(menu.Mainmenu())
	W.SetContent(container.NewScroll(
		container.NewVBox(
			MainContainer.SearchContainer(), MainContainer.CardContainer(),
		)))

	W.ShowAndRun()
}

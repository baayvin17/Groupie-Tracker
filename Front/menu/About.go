package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main/res"
	"strings"
)

func About() func() {
	return func() {
		w := res.CreateNewWindow(res.L.MenuHelp[0])

		MainContainer := container.NewVBox()
		MainContainer.Add(widget.NewLabel(res.L.AboutAppName + " " + res.L.AppName))
		MainContainer.Add(widget.NewLabel(res.L.AboutAppVersion + " " + res.L.AppVersion))
		MainContainer.Add(widget.NewLabel(res.L.AboutAppDescription + " " + res.L.AppDescription))
		MainContainer.Add(widget.NewLabel(res.L.AboutAppAuthor + " " + res.L.AppAuthor))

		MainContainer.Add(layout.NewSpacer())

		MainContainer.Add(container.NewHBox(
			widget.NewLabel(res.L.AboutAppRepository),
			widget.NewButtonWithIcon(GetRepositoryName(res.UrlRepository), res.AboutGiteaIcon, func() {
				res.OpenBrowser(res.UrlRepository)
			}),
		))

		w.SetContent(MainContainer)

		w.Resize(fyne.NewSize(400, 250))
		w.SetFixedSize(true)
		w.Show()
	}
}

func GetRepositoryName(url string) string {
	//remove https://
	url = url[8:]
	//remove everything before the last /
	url = url[len(url)-len(url)+strings.LastIndex(url, "/")+1:]
	return url
}

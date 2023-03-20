package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"main/res"
	"main/res/save"
	"os"
	"sort"
)

var (
	tmp = save.PREF.Language
)

func language() *fyne.Container {
	Hbox := container.NewHBox()
	//Add label
	Hbox.Add(
		widget.NewLabel(res.L.PreferencesLanguages),
	)
	//Add combobox
	Hbox.Add(languageChoice())

	return Hbox
}

func Preferences() func() {
	return func() {
		w := res.CreateNewWindow(res.L.MenuFile[0])

		MainContainer := container.NewVBox()
		MainContainer.Add(language())
		MainContainer.Add(theme())
		MainContainer.Add(container.NewHBox(layout.NewSpacer(), widget.NewButton(res.L.BUTTONS["Save"], func() {
			//if tmp != save.PREF.Language {
			//	quit = true
			//}
			if tmp != save.PREF.Language {
				os.Exit(0)
			} else {
				w.Close()
			}
		})))
		w.SetContent(MainContainer)

		w.Show()
	}

}

func languageChoice() *widget.Select {
	//create a new empty array
	var lang []string
	//for all file in the lang folder add it to the array
	files, err := os.ReadDir(res.LangPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		title := file.Name()[0 : len(file.Name())-5]
		lang = append(lang, title)
	}

	//create a new combobox with the array
	combobox := widget.NewSelect(lang, func(s string) {
		res.InitLang(res.ThemePath + s + ".json")
		tmp = save.PREF.Language
		save.PREF.Language = s
		res.CreateFile(save.PrefPath, save.PREF)
		//os.Exit(0)
	})

	//log.Printf("Language: %s", combobox.Options)
	//print(save.PREF.Language)
	combobox.PlaceHolder = save.PREF.Language
	combobox.Refresh()
	return combobox
}

func theme() *fyne.Container {
	Hbox := container.NewHBox()
	//Add label
	Hbox.Add(
		widget.NewLabel(res.L.PreferencesTheme),
	)
	//Add combobox
	Hbox.Add(themeChoice())

	return Hbox
}

func themeChoice() fyne.CanvasObject {
	//create a new empty array
	var theme []string
	//for all file in the lang folder add it to the array
	files, err := os.ReadDir(res.ThemePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		title := file.Name()[0 : len(file.Name())-5]
		theme = append(theme, title)
	}

	//create a new combobox with the array
	sort.Strings(theme)

	combobox := widget.NewSelect(theme, func(s string) {

		res.InitTheme("res/theme/" + s + ".json")
		save.PREF.Theme = s

		res.CreateFile(save.PrefPath, save.PREF)

		//os.Exit(0)
	})
	combobox.SetSelected(save.PREF.Theme)
	return combobox
}

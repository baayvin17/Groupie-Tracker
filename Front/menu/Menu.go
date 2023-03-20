package menu

import (
	"fyne.io/fyne/v2"
	"main/res"
)

func Mainmenu() *fyne.MainMenu {
	// File Menu items
	FmenuItem1 := fyne.NewMenuItem(res.L.MenuFile[0], Preferences())
	// Search Menu Items
	SmenuItem1 := fyne.NewMenuItem(res.L.MenuSearch[0], nil)
	SmenuItem2 := fyne.NewMenuItem(res.L.MenuSearch[1], nil)
	SmenuItem3 := fyne.NewMenuItem(res.L.MenuSearch[2], nil)
	SmenuItem4 := fyne.NewMenuItem(res.L.MenuSearch[3], nil)
	// Help Menu Items
	HmenuItem1 := fyne.NewMenuItem(res.L.MenuHelp[0], About())
	// New Menu
	FileMenu := fyne.NewMenu(res.L.MenuMain[0], FmenuItem1)
	SearchMenu := fyne.NewMenu(res.L.MenuMain[1], SmenuItem1, SmenuItem2, SmenuItem3, SmenuItem4)
	HelpMenu := fyne.NewMenu(res.L.MenuMain[2], HmenuItem1)
	// New main menu
	menu := fyne.NewMainMenu(FileMenu, SearchMenu, HelpMenu)
	// setup main menu
	return menu
}

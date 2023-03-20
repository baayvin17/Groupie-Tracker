package MainContainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"main/Front/Search_Bar"
)

func CardContainer() *fyne.Container {
	// New container
	MainContainer := container.NewAdaptiveGrid(1)
	MainContainer.Add(Cards())
	return MainContainer
}
func SearchContainer() *fyne.Container {
	SearchContainer := container.NewVBox(
		Search_Bar.FiltreCheckBox1(),
		Search_Bar.FiltreCheckBox2(),
		Search_Bar.FiltreMinMax1(),
		Search_Bar.FiltreMinMax2(),
		Search_Bar.FiltreMinMaxSlider(),
		Search_Bar.FiltreLoc(),
		Search_Bar.SearchBar(),
		Search_Bar.SearchButton(),
		Search_Bar.SearchResult(),
	)

	return SearchContainer
}

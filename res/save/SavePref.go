package save

import (
	"main/res"
)

type Preferences struct {
	Language string `json:"Language"`
	Theme    string `json:"Theme"`
}

var (
	PrefPath = "res/save/preferences.json"
	PREF     = new(Preferences)
)

func InitPref() {
	res.LoadFile(PrefPath, PREF, PrefPath)
	res.InitLang(res.LangPath + PREF.Language + ".json")
	res.InitTheme(res.ThemePath + PREF.Theme + ".json")
}

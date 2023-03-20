package res

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type Lang struct {
	AppName        string `json:"APP_NAME"`
	AppVersion     string `json:"APP_VERSION"`
	AppAuthor      string `json:"APP_AUTHOR"`
	AppDescription string `json:"APP_DESCRIPTION"`

	MenuMain   []string `json:"MENU_MAIN"`
	MenuFile   []string `json:"MENU_FILE"`
	MenuSearch []string `json:"MENU_SEARCH"`
	MenuHelp   []string `json:"MENU_HELP"`

	BUTTONS map[string]string `json:"BUTTONS"`

	SearchBar string `json:"SEARCH_BAR"`

	PreferencesLanguages string `json:"PREFERENCES_LANGUAGES"`
	PreferencesTheme     string `json:"PREFERENCES_THEME"`

	AboutAppName        string `json:"ABOUT_APP_NAME"`
	AboutAppVersion     string `json:"ABOUT_APP_VERSION"`
	AboutAppAuthor      string `json:"ABOUT_APP_AUTHOR"`
	AboutAppDescription string `json:"ABOUT_APP_DESCRIPTION"`
	AboutAppRepository  string `json:"ABOUT_APP_REPOSITORY"`

	SearchCategoryArtiste      string `json:"SEARCH_CATEGORY_ARTISTE"`
	SearchCategoryMember       string `json:"SEARCH_CATEGORY_MEMBER"`
	SearchCategoryLocation     string `json:"SEARCH_CATEGORY_LOCATION"`
	SearchCategoryConcertDate  string `json:"SEARCH_CATEGORY_CONCERT_DATE"`
	SearchCategoryCreationDate string `json:"SEARCH_CATEGORY_CREATION_DATE"`
	SearchCategoryAlbumDate    string `json:"SEARCH_CATEGORY_ALBUM_DATE"`

	SearchResultPlaceholder   string `json:"SEARCH_RESULTAT_PLACEHORDER"`
	SearchNoResultPlaceholder string `json:"SEARCH_NO_RESULTAT_PLACEHORDER"`

	ArtistFirstAlbum string `json:"ARTIST_FIRST_ALBUM"`
	ArtistConcerts   string `json:"ARTIST_CONCERTS"`

	FilterBand         string `json:"FILTRE_BAND"`
	FilterMember       string `json:"FILTRE_MEMBER"`
	FilterLocation     string `json:"FILTRE_LOCATION"`
	FilterConcertDate  string `json:"FILTRE_CONCERT_DATE"`
	FilterCreationDate string `json:"FILTRE_CREATION_DATE"`
	FilterAlbumDate    string `json:"FILTRE_ALBUM_DATE"`

	FilterNbMember        string `json:"FILTRE_NB_MEMBER"`
	FilterMaxCreationDate string `json:"FILTRE_MAX_CREATION_DATE"`
	FilterMinCreationDate string `json:"FILTRE_MIN_CREATION_DATE"`
	FilterMaxConcertDate  string `json:"FILTRE_MAX_CONCERT_DATE"`
	FilterMinConcertDate  string `json:"FILTRE_MIN_CONCERT_DATE"`
	FilterSelectLocation  string `json:"FILTRE_SELECT_LOCATION"`
}

type Theme struct {
	ColorBackground      color.NRGBA `json:"COLOR_BACKGROUND"`
	ColorForeground      color.NRGBA `json:"COLOR_FOREGROUND"`
	ColorShadow          color.NRGBA `json:"COLOR_SHADOW"`
	ColorScrollbar       color.NRGBA `json:"COLOR_SCROLLBAR"`
	ColorHover           color.NRGBA `json:"COLOR_HOVER"`
	ColorFocus           color.NRGBA `json:"COLOR_FOCUS"`
	ColorText            color.NRGBA `json:"COLOR_TEXT"`
	ColorDisabled        color.NRGBA `json:"COLOR_DISABLED"`
	ColorPlaceholder     color.NRGBA `json:"COLOR_PLACEHOLDER"`
	ColorInputBackground color.NRGBA `json:"COLOR_INPUT_BACKGROUND"`
	ColorButton          color.NRGBA `json:"COLOR_BUTTON"`
	ColorPrimary         color.NRGBA `json:"COLOR_PRIMARY"`
	ColorPressed         color.NRGBA `json:"COLOR_PRESSED"`
}

var (
	// LangPath --------------------------------------------------------------------------------------------------------
	LangPath = "res/lang/"
	L        = new(Lang)

	// ThemePath -------------------------------------------------------------------------------------------------------
	ThemePath = "res/theme/"
	T         = new(Theme)

	// Application icon
	AppIcon, _ = fyne.LoadResourceFromURLString(UrlAppIcon)

	// URLS ------------------------------------------------------------------------------------------------------------

	UrlApi         = "https://groupietrackers.herokuapp.com/api/"
	UrlApiArtists  = UrlApi + "artists"
	UrlApiRelation = UrlApi + "relation"

	UrlRepository = "https://ytrack.learn.ynov.com/git/walberic/GroopieTracker"

	UrlWikipedia = "https://fr.wikipedia.org/wiki/"

	UrlAppIcon = "https://www.pngkey.com/png/full/356-3562324_app-logo-youtube-music-icon.png"

	// UrlAboutGiteaIcon ----------------------------------------------------------------------------------------------
	UrlAboutGiteaIcon = "https://upload.wikimedia.org/wikipedia/commons/b/bb/Gitea_Logo.svg"
	AboutGiteaIcon, _ = fyne.LoadResourceFromURLString(UrlAboutGiteaIcon)
)

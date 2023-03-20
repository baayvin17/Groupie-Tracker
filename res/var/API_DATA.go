package _var

type relations struct {
	Index []relation `json:"index"`
}

type relation struct {
	ID           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"` //ville ; [date]
}

type artistes struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var (
	RData relations
	AData []artistes
)

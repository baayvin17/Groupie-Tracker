package res

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func InitLang(path string) {
	LoadFile(path, &L, LangPath+"en.json")
}

func CreateNewWindow(title string) fyne.Window {
	w := fyne.CurrentApp().NewWindow(title)
	w.Resize(fyne.NewSize(800, 600))

	//set the icon
	w.SetIcon(AppIcon)
	return w
}

func LoadFile(path string, data interface{}, defaultPath string) {
	file, err := os.Open(path)
	if err != nil {
		file, err = os.Open(defaultPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//read file and put it in a byte array
	byteValue, _ := io.ReadAll(file)

	err = json.Unmarshal(byteValue, data)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(path string, data interface{}) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//read file and put it in a byte array
	byteValue, _ := json.Marshal(data)

	_, err = file.Write(byteValue)
	if err != nil {
		log.Fatal(err)
	}
}

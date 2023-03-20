package res

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type MyTheme struct{}

func (MyTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return T.ColorBackground
	case theme.ColorNameButton:
		return T.ColorButton
	case theme.ColorNameDisabledButton:
		return T.ColorDisabled
	case theme.ColorNameDisabled:
		return T.ColorDisabled
	case theme.ColorNameError:
		return T.ColorDisabled
	case theme.ColorNameFocus:
		return T.ColorFocus
	case theme.ColorNameForeground:
		return T.ColorForeground
	case theme.ColorNameHover:
		return T.ColorHover
	case theme.ColorNameInputBackground:
		return T.ColorInputBackground
	case theme.ColorNamePlaceHolder:
		return T.ColorPlaceholder
	case theme.ColorNamePrimary:
		return T.ColorPrimary
	case theme.ColorNamePressed:
		return T.ColorPressed
	case theme.ColorNameScrollBar:
		return T.ColorScrollbar
	case theme.ColorNameShadow:
		return T.ColorShadow
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (MyTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}

func InitTheme(path string) {
	if path == "res/theme/default.json" {
		fyne.CurrentApp().Settings().SetTheme(theme.DefaultTheme())
	} else {
		LoadFile(path, &T, ThemePath+"default.json")
		fyne.CurrentApp().Settings().SetTheme(MyTheme{})
	}
}

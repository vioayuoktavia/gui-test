package data

import (
	"fyne.io/fyne/v2"
)

type menuLayout struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	menuLayouts = map[string]menuLayout{
		"welcome": {"Welcome", "", welcomeScreen},
		"canvas": {"Canvas",
			"See the canvas capabilities.",
			canvasScreen,
		},
	}
	// layoutIndex  defines how our data should be laid out in the index tree
	layoutIndex = map[string][]string{
		"": {"welcome", "canvas"},
	}
)

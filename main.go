package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentLayout = "currentLayout"

var topWindow fyne.Window

func main() {
	apps := app.NewWithID("two.demo")
	windows := apps.NewWindow("TWO Demo")

	topWindow = windows

	windows.SetMainMenu(makeMenu(apps, windows))
	windows.SetMaster()
	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setLayout := func(d data.menuLayout) {
		title.SetText(d.Title)
		intro.SetText(d.Intro)

		content.Objects = []fyne.CanvasObject{d.View(windows)}
		content.Refresh()
	}

	layout := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(setLayout, true), layout)
	split.Offset = 0.2
	windows.SetContent(split)

	windows.Resize(fyne.NewSize(640, 460))
	windows.ShowAndRun()

}

func makeMenu(apps fyne.App, windows fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("Hi", nil)
	file := fyne.NewMenu("File", newItem)

	return fyne.NewMainMenu(file)
}

func makeNav(setLayout func(layout data.menuLayout), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return data.layoutIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := data.layoutIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := data.menuLayouts[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := data.menuLayouts[uid]; ok {
				a.Preferences().SetString(preferenceCurrentLayout, uid)
				setLayout(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentLayout, "welcome")
		tree.Select(currentPref)
	}

	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

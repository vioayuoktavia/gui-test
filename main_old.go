package main

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println(data)
	a := app.New()
	w := a.NewWindow("TWO Automation Monitoring")
	list_of_log := []string{""}
	time_now := time.Now()
	time_formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		time_now.Year(), time_now.Month(), time_now.Day(), time_now.Hour(), time_now.Minute(), time_now.Second())

	log_monitoring := widget.NewLabel("")

	welcome_layout := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		widget.NewLabel("Welcome, team!"),
		layout.NewSpacer(),
	)

	btn_start := widget.NewButton("Start", func() {
		template_log := time_formatted + " - " + "Started RPA" + "\n"
		list_of_log = append(list_of_log, template_log)
		logs := strings.Join(list_of_log, " ")
		log_monitoring.SetText(logs)

		//drop here your func to call start RPA
	})

	btn_stop := widget.NewButton("Stop", func() {
		template_log_stop := time_formatted + " - " + "Stoping RPA" + "\n"
		list_of_log = append(list_of_log, template_log_stop)
		logs_stop := strings.Join(list_of_log, " ")
		log_monitoring.SetText(logs_stop)

		//drop here your func to call stop RPA
	})

	btn_color_red := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})

	btn_color_green := canvas.NewRectangle(
		color.NRGBA{R: 0, G: 255, B: 0, A: 255})

	btn_start_container := container.New(
		layout.NewMaxLayout(),
		btn_color_green,
		btn_start,
	)

	btn_stop_container := container.New(
		layout.NewMaxLayout(),
		btn_color_red,
		btn_stop,
	)

	button_layout := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		btn_start_container,
		btn_stop_container,
		widget.NewButton("Clear Log", func() {
			list_of_log = []string{""}
			log_monitoring.SetText("")
		}),

		layout.NewSpacer(),
	)

	w.SetContent(container.New(layout.NewVBoxLayout(), welcome_layout, button_layout, log_monitoring))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

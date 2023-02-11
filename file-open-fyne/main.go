package main

// import fyne
import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("davids little file converter tool")
	w.Resize(fyne.NewSize(800, 400))
	btn := widget.NewButton("Please select the .txt file", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				// read files
				data, _ := ioutil.ReadAll(r)
				// reader will read file and store data
				// now result
				result := fyne.NewStaticResource("name", data)
				// lets display our data in label or entry
				entry := widget.NewMultiLineEntry()
				// string() function convert byte to string
				entry.SetText(string(result.StaticContent))
				// Lets show and setup content
				// tile of our new window
				w := fyne.CurrentApp().NewWindow(
					string(result.StaticName)) // title/name
				w.SetContent(container.NewScroll(entry))
				w.Resize(fyne.NewSize(400, 400))
				// show/display content
				w.Show()
				// we are almost done
			}, w)
		// fiter to open .txt files only
		// array/slice of strings/extensions
		file_Dialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()
		// Show file selection dialog.
	})
	// lets show button in parent window
	w.SetContent(container.NewVBox(
		btn,
	))
	w.ShowAndRun()
}

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewFillObject(title string) *widget.Label {
	label := widget.NewLabel(title)
	return label
}

type EntryLabel struct {
	text string
	entry *widget.Entry
	label *widget.Label
	isEntry bool
}

func NewEntryLabel(text string) (e*EntryLabel) {
	e.text = text
	e.label = widget.NewLabel(text)
	e.entry = widget.NewEntry()
	e.entry.OnChanged = func (s string)  {
		e.entry.Text = s
		e.text = s
	}
	e.entry.Text = text
	e.isEntry = true
	return e
}

func (e *EntryLabel) Build() fyne.CanvasObject {
	cont := container.NewCenter(e.label,e.entry)
	return cont
}

// Swap between the entry widget and the label widget
func (e *EntryLabel) Swap() {
	e.isEntry = !e.isEntry
	if e.isEntry {
		e.entry.Show()
		e.label.Hide()
	} else {
		e.entry.Hide()
		e.label.Show()
	}
	
}
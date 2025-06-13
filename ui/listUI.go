package ui

import (
	"rinogodson/DreamShell/filehandler"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ListUI() {
	app := tview.NewApplication()

	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle("[#D787FF]╯✨[#FFD8FF] Dreams [#D787FF]✨[#D787FF]╰")
	list.SetTitleColor(tcell.Color225)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlJ {
			list.SetCurrentItem(list.GetCurrentItem() + 1)
		} else if event.Key() == tcell.KeyCtrlK {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		}
		return event
	})

	dreams := filehandler.GetFiles()

	for _, dream := range dreams {
		item := filehandler.ParseInput(dream.Name())
		list.AddItem(item[0], item[1], 0, nil)
	}

	listBox := tview.NewFlex()
	listBox.SetTitle("[#D787FF]╯✨[#FFD8FF] DreamShell ✨[#D787FF]╰").SetBorder(true)
	listBox.SetBorderColor(tcell.Color177)
	listBox.SetTitleColor(tcell.Color225)
	listBox.SetBorderPadding(1, 0, 3, 3)

	listBox.AddItem(list, 0, 1, true)

	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	container := tview.NewPages()
	container.AddPage("main", modal(listBox, 70, 35), true, true)
	if err := app.SetRoot(container,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

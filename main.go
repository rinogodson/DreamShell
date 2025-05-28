package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	tview.Borders.Horizontal = '─'
	tview.Borders.Vertical = '│'
	tview.Borders.TopLeft = '╭'
	tview.Borders.TopRight = '╮'
	tview.Borders.BottomLeft = '╰'
	tview.Borders.BottomRight = '╯'

	tview.Borders.BottomRightFocus = '┛'
	tview.Borders.BottomLeftFocus = '┗'
	tview.Borders.TopRightFocus = '┓'
	tview.Borders.TopLeftFocus = '┏'
	tview.Borders.HorizontalFocus = '━'
	tview.Borders.VerticalFocus = '┃'

	themeColor := tcell.Color147
	app := tview.NewApplication()

	titleArea := tview.NewTextArea().SetPlaceholder("Dream Title...")
	titleArea.SetTitle("╮ Title ╭─").SetBorder(true)
	titleArea.SetBorderColor(themeColor)
	titleArea.SetTitleColor(themeColor)
	titleArea.SetBorderPadding(0, 0, 1, 1)
	titleArea.SetTextStyle(tcell.StyleDefault.Foreground(tcell.Color230).Bold(true))
	titleArea.SetTitleAlign(tview.AlignRight)
	titleArea.SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorSilver))
	
	bodyArea := tview.NewTextArea().SetPlaceholder("I dreamed about...")
	bodyArea.SetTitle("╮ Body ╭─").SetBorder(true)
	bodyArea.SetBorderColor(themeColor)
	bodyArea.SetTitleColor(themeColor)
	bodyArea.SetBorderPadding(0, 0, 1, 1)
	bodyArea.SetTextStyle(tcell.StyleDefault.Foreground(tcell.Color230))
	bodyArea.SetTitleAlign(tview.AlignRight)
	bodyArea.SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorSilver))

	tagArea := tview.NewTextArea().SetPlaceholder("Tags example: #funny #sad #nightmare")
	tagArea.SetTitle("╮ Tags (optional) add #lucid if it was lucid ╭─").SetBorder(true)
	tagArea.SetTitleAlign(tview.AlignRight)
	tagArea.SetBorderColor(themeColor)
	tagArea.SetTitleColor(themeColor)
	tagArea.SetBorderPadding(0, 0, 1, 1)
	tagArea.SetTextStyle(tcell.StyleDefault.Foreground(tcell.Color193))
	tagArea.SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorSilver))

//	lucidToggle := tview.NewCheckbox().SetLabel("dkj").SetTitle("╮ Lucid? ╭─").SetBorder(true)

	mainView := tview.NewGrid().SetRows(3, 0, 3).SetColumns(0)
	mainView.AddItem(titleArea,   0, 0, 1, 1, 0, 0, true)
	mainView.AddItem(bodyArea,    1, 0, 1, 1, 0, 0, false)
	mainView.AddItem(tagArea,     2, 0, 1, 1, 0, 0, false)
//	mainView.AddItem(lucidToggle, 2, 1, 3, 1, 0, 0, false)
	mainView.SetTitle("╯✨ DreamShell ✨╰").SetBorder(true)
	mainView.SetBorderColor(tcell.Color177)
	mainView.SetTitleColor(tcell.Color225)
	mainView.SetBorderPadding(1, 1, 3, 3)

	container := tview.NewFlex()
	container.SetDirection(tview.FlexRow)
	container.AddItem(mainView, 0, 1, true)
	_, _, width, _ := container.GetInnerRect()

	container.SetBorderPadding(1, 1, getPadding(width), getPadding(width))
	if err := app.SetRoot(container,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func getPadding(width int) int {
	if width > 15 {
		return 15
	}
	return 0
}

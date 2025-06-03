package main

import (
	"rinogodson/DreamShell/filehandler"
	"rinogodson/DreamShell/ui"
)

func main()  {
	ui.TviewConfigInit()
	ui.NewUI()
	filehandler.CreateFile("test")
}

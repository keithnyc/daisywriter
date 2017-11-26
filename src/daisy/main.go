package main

import (
	"log"
	"os"
	"github.com/jroimartin/gocui"
)

// Global Key Configuration
var keys []Key = []Key{
	Key{"", gocui.KeyCtrlQ, actionGlobalQuit},
	Key{"", gocui.KeyCtrlO, actionLoadFile},
	Key{"", gocui.KeyCtrlS, actionSaveFile},
	Key{"", gocui.KeyArrowUp, actionCursorUp},
	Key{"", gocui.KeyArrowDown, actionCursorDown},
	Key{"", gocui.KeyCtrlH, actionBeginningOfDocument},
	Key{"", gocui.KeyPgdn, actionPgDn},
	Key{"", gocui.KeyPgup, actionPgUp},
	Key{"", gocui.KeyEnd, actionEnd},
	Key{"", gocui.KeyHome, actionHome},
	Key{"", gocui.KeyTab, actionTab},
}

var currentFilePath = ""
var isFileModified =false

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	
	g.SetManagerFunc(uiLayout)
	
	if err := uiKey(g); err != nil{
		log.Panicln(err)
	}
	
	//check if file was passed as arg
	if len(os.Args) > 1{
		
		filePath := os.Args[1]
		
		//create the file if it doesn't exist
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
		  f, _ :=os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		  f.Close()
		}
		
		maxX, maxY := g.Size()
		viewEditor(g, maxX, maxY)
		loadViewEditor(g, filePath)
		
	}
	
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	
	
}

func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	viewTitle(g, maxX, maxY)
	viewEditor(g, maxX, maxY)
	viewStatusBar(g, maxX, maxY)
	updateUI(g)
	return nil
}

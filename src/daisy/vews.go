package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/jroimartin/gocui"
	"github.com/willf/pad"
)
const AUTHOR = "by Keith(keith@unixloft.com)"
const APP = "▒▒ Daisy Writer ≡"

var MENU_DISPLAYING bool = false



func editorAuthor() string {
	return fmt.Sprintf(" %s ", AUTHOR)
}

func editorTitleBanner() string {
	return fmt.Sprintf(" %s", APP)
}

func updateViewTitle(g *gocui.Gui, title string) error{
	g.DeleteView("title")
	lMaxX, _ := g.Size()
	if v, err := g.SetView("title", -1, -1, lMaxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		
		// Settings
		v.Frame = false
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse
		
		title := editorTitleBanner() + pad.Left(" »" + title, lMaxX-len(editorTitleBanner()), "≡")
		fmt.Fprintln(v, title)
	}
	return nil
}
// View: Title bar
func viewTitle(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("title", -1, -1, lMaxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse
		
		
		title := editorTitleBanner() + pad.Left(editorAuthor(), lMaxX-len(editorTitleBanner()), " ")
		// Content
		fmt.Fprintln(v, title)
	}

	return nil
}

func loadViewEditor(g *gocui.Gui, filePath string) error {
	v, err := g.View("main")
	v.Clear()
	
	currentFilePath = filePath
	
	if err != nil {
			panic(err.Error())
	}
	b, err := ioutil.ReadFile(filePath)
	
	if err != nil{
		panic(err.Error())
	}
	v.Autoscroll = false
	fmt.Fprintf(v, "%s", b)
	
	updateViewTitle(g, filePath)
	return nil	

}

// View: Text Editor
func viewEditor(g *gocui.Gui, lMaxX int, lMaxY int) error{
	
	if v, err := g.SetView("main", 0, 1, lMaxX-1, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true
		v.Frame = false
		v.Autoscroll = true
		
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}


// View: Status bar
func viewStatusBar(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("status", -1, lMaxY-2, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
		v.BgColor = gocui.ColorBlack
		v.FgColor = gocui.ColorWhite
		
				
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse

		// Content
		statusBarText(g)
	}

	return nil
}

// Status Bar
func statusBarText(g *gocui.Gui)  {
	//lMaxX, _ := g.Size()
	v, err := g.View("status")
	editView, err := g.View("main")
	
	if err != nil {
		return
	}
	
	
	cx, cy := editView.Cursor()
	
	var str = ""
	str = fmt.Sprintf("| col:%d | row:%d | lines:%d |",
							 cx+1, cy+1, 
							 strings.Count(editView.Buffer(), "\n")) 
	
	v.Clear()
	i := 5
	b := " "
	b = b + frameText("^S")+ " SAVE   "
	b = b + frameText("^Q") +" QUIT   "
	b = b + frameText("^W") +" WRAP   "
	
	
	var curText = pad.Left(str, len(str) - len(b), "X") 
	var statText = pad.Left(b, i, " ")+ " " + curText
	
	fmt.Fprintln(v, statText)
	
	return
}


package main

import (
	"strings"
	"github.com/jroimartin/gocui"
)


func IsAtBottom(g *gocui.Gui) bool {    
    v, _ := g.View("main")
    _, cy := v.Cursor()
    
    
    if cy >= (strings.Count(v.ViewBuffer(), "\n") -1) {        
        return true
    } else {       
    	
        return false
    }
	
	
}


func IsAtVisualBottom(g *gocui.Gui) bool {    
    v, _ := g.View("main")
    _, cy := v.Cursor()
    if cy >= (strings.Count(v.Buffer(), "\n") - 1) {        
        return true
    } else {        
        return false
    }

}
// Get view line (relative to the cursor)
func getViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	return l, err
}

// Get the next view line (relative to the cursor)
func getNextViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy + 1); err != nil {
		l = ""
	}

	return l, err
}

// Set view cursor to line
func setViewCursorToLine(g *gocui.Gui, v *gocui.View, lines []string, selLine string) error {
	ox, _ := v.Origin()
	cx, _ := v.Cursor()
	for y, line := range lines {
		if line == selLine {
			if err := v.SetCursor(ox, y); err != nil {
				if err := v.SetOrigin(cx, y); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

package main

import (
	"io"
	"io/ioutil"
	"strings"
	"github.com/jroimartin/gocui"
)

func actionTab(g *gocui.Gui, v *gocui.View) error {
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	return nil
}
func actionEnd(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	linestr, _ := v.Line(cy)
	v.SetCursor(len(linestr), cy)
	//v.MoveCursor(len(linestr), cy, false)
	return nil
}

func actionHome(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	v.SetCursor(0, cy)
	return nil
}

func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func actionLoadFile(g *gocui.Gui, v *gocui.View) error{
	//loadViewEditor(g, "/home/keith/test.txt")
	return nil
}

func actionPgUp(g *gocui.Gui, v *gocui.View) error{
	if v != nil {
		_, oy := v.Origin()
		cx, cy := v.Cursor()
		
		if cy <= 10 || oy <= 10{
			actionBeginningOfDocument(g, v)
			return nil
		}
		
		if err := v.SetCursor(cx, cy-10); err != nil && oy > 0 {
			if err = v.SetOrigin(cx, cy-10); err != nil {
				return nil
			}
		}
		
	}else{
		panic("No view found")
		
	}
	return nil
}

func actionPgDn(g *gocui.Gui, v *gocui.View) error{
	
	if(IsAtBottom(g)){
		return nil
	} 
	
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		
		if (cy+10) >= (strings.Count(v.ViewBuffer(), "\n") -1){
			v.SetCursor(cx, strings.Count(v.ViewBuffer(), "\n") )
		}else{
			if err := v.SetCursor(cx, cy+10); err != nil {
				if err := v.SetOrigin(ox, oy+10); err != nil {
					return nil
				}
			}	
		}
		
		
	}else{
		panic("No view found")
		
	}
	return nil
}
func actionBeginningOfDocument(g *gocui.Gui, v *gocui.View) error{
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	actionScroll(g, 0)
	return nil
}

func actionScroll(g *gocui.Gui, dy int) {
    // Grab the view that we want to scroll.
    v, _ := g.View("main")

    // Get the size and position of the view.    
    ox, oy := v.Origin()
	
	if IsAtBottom(g){
		v.Autoscroll = true
	}else{
		v.Autoscroll = false
        v.SetOrigin(ox, oy+dy)
	}   
}

func actionCursorUp(g *gocui.Gui, v *gocui.View) error { 	
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
			actionScroll(g, cy)
		}
		
	}else{
		panic("No view found")
		
	}
	
	actionHome(g, v)
	return nil
}

func updateUI(g *gocui.Gui){
	g.Update(func(g *gocui.Gui) error {
				statusBarText(g)
				return nil
	})
}
func actionCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return nil
			}
		}
	}
	return nil
}


func actionSaveFile(g *gocui.Gui, v *gocui.View) error{
	
	p := make([]byte, len(v.Buffer()))
	v.Rewind()
	
	for {
		n, err := v.Read(p)
		if n > 0 {
			if err := ioutil.WriteFile(currentFilePath, p[:n], 655) ; err != nil {
				panic(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	return nil
}
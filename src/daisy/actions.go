package main

import (
	"io"
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/jroimartin/gocui"
)

func actionDeleteLine(g *gocui.Gui, v *gocui.View) error{
	_, cy := v.Cursor()
	maxX, _ := v.Size()
	
	actionEnd(g,v)
	str, _ := v.Line(cy)
	
	i:=0
	lstr := len(str)
	fmt.Println(lstr)
	if !v.Wrap{
		for i < maxX && i <= lstr{
			v.EditDelete(true)
			i++
		}
	}else{
		for i <= lstr{
			v.EditDelete(true)	
			i++
		}	
	}
	
	return nil
}

func actionToggleWrap(g *gocui.Gui, v *gocui.View) error{
	v.Wrap = !(v.Wrap)
	return nil
}
func actionTab(g *gocui.Gui, v *gocui.View) error {
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	return nil
}
func actionEnd(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := v.Size()
	cx, cy := v.Cursor()
	linestr, _ := v.Line(cy)
	
	lstr := len(linestr)
	
	if len(linestr) > maxX{
		if v.Wrap{
			actionHome(g,v)
			wraplines := (lstr / maxX)
			
			if wraplines < 2{
				wraplines = 2
			}
			
			fmt.Println(wraplines)
			origPos := cy
			cy = ((cy + wraplines)) - (origPos - wraplines) - 1
			lbxpos := (lstr % maxX) - 2
			cx = lbxpos
		}else{
			cx = maxX -1
		}
	}else{
		cx = lstr
	} 
	
	v.SetCursor(cx, cy)
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
	
	actionHome(g, v)
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
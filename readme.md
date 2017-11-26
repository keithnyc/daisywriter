Daisy Writer is a fast, efficient terminal editor written using Go. 

![Image of Daisy Writer](https://i.imgur.com/cg6dMkJ.png)

### Usage ###

Daisy Writer doesn't have a full menu system yet, so make sure you open existing files
by passing the file name from the command line (e.g. daisy mydoc.txt). If the file doesn't
exist, it will be created.

#### Key Bindings ####

Action              | Key Combination
------------------- | ---------------
Save                | Ctrl+S
Quit (immediate)    | Ctrl+Q
Toggle Line Wrap    | Ctrl+W


Note: Make sure your terminal emulator is set to unicode.

### Work In Progress ###

Daisy Writer is still in development and not really ready for prime time. 

If you lose all your work or your thesis you spent 80 hours writing explodes,
I'm not responsible :)

Pull requests are appreciated. 

### Inspiration ###

I'm not a big fan of nano and while I love vim, I wanted something a little more akin
to a word processor than what vim offers. I really love wordgrinder but just find it a bit
too buggy on FreeBSD, but that was really the closest thing I was aiming for.

The name Daisy Writer is named for my persian cat. She watches me write and she's really
small, yet full featured. So yeah, seemed like a good idea.

## Thanks ##

A big thank you to antirez for his most excellent tutorial on writing your own editor in C. His 
tutorial was super helpful and well put together.

https://viewsourcecode.org/snaptoken/kilo/index.html

Also thanks to the gocui project and Julien Breux's pody project whose code was really invaluable in 
understanding gocui.

gocui:
https://github.com/jroimartin/gocui

Pody:
https://github.com/JulienBreux/pody





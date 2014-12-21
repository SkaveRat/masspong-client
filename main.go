package main

import (
	termbox "github.com/nsf/termbox-go"
	"fmt"
	"net"
)

func main() {

	termbox.Init()

	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		panic("could not connect")
	}
	mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				fmt.Fprintf(conn, "1 up\n")
			case termbox.KeyArrowDown:
				fmt.Fprintf(conn, "1 down\n")
			case termbox.KeyEsc:
				break mainloop
			}

		}
	}
}

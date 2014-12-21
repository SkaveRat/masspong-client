package main

import (
	termbox "github.com/nsf/termbox-go"
	"fmt"
	"net"
	"flag"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
}

func main() {

	var port int
	var ip string
	var player int

	flag.IntVar(&port, "port", 1337, "remote port")
	flag.StringVar(&ip, "host", "localhost", "remote host")
	flag.IntVar(&player, "player", 1, "player")

	flag.Parse()

	remote := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println(remote)

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		panic("could not connect")
	}

	termbox.Init()
	defer termbox.Close()
	mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				fmt.Fprintf(conn, fmt.Sprintf("%d up\n", player))
			case termbox.KeyArrowDown:
				fmt.Fprintf(conn, fmt.Sprintf("%d down\n", player))
			case termbox.KeyEsc:
				break mainloop
			}

		}
	}
}

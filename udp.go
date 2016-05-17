package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	laddr = flag.String("l", "0.0.0.0:8060", "local address")
)

func Start() {
	c, err := net.ListenPacket("udp", *laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	log.Printf("%v\n", c.LocalAddr())
	for {
		if err := c.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			log.Fatal(err)
		}

		rb := make([]byte, 2000)

		n, addr, err := c.ReadFrom(rb)
		if err != nil {
			log.Print(err)
		}
		log.Printf("%d %v\n", n, addr)
	}
}

func main() {
	flag.Parse()
	Start()
}

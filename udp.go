package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	laddr = flag.String("l", "0.0.0.0:8060", "local address")
	raddr = flag.String("r", "", "remote address")
)

func heartbeat(c net.PacketConn, addr net.Addr) {
	const beatInterval = 3 * time.Second
	ticker := time.NewTicker(beatInterval)
	go func() {
		for {
			<-ticker.C
			n, err := c.WriteTo([]byte("zjk"), addr)
			if err != nil {
				log.Print(err)
			} else {
				log.Printf("send %d\n", n)
			}
		}
	}()
}

func Start() {
	c, err := net.ListenPacket("udp4", *laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if len(*raddr) > 0 {
		paddr, err := net.ResolveUDPAddr("udp", *raddr)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("remote: %v\n", *paddr)
		heartbeat(c, paddr)
	}

	log.Printf("%v\n", c.LocalAddr())
	for {
		if err := c.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
			log.Fatal(err)
		}

		rb := make([]byte, 2000)

		n, addr, err := c.ReadFrom(rb)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("%d %v\n", n, addr)
		}
	}
}

func main() {
	flag.Parse()
	Start()
}

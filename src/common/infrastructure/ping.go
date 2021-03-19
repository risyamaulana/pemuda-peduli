package infrastructure

import (
	"log"
	"net"
	"time"
)

// Ping check if service alive / up
func Ping(host string) {
	log.Println("Checking if service alive")

	for {
		log.Printf("Ping %s ...", host)
		conn, err := net.Dial("tcp", host)
		if err == nil {
			_ = conn.Close()
			log.Printf("%s is alive!", host)
			return
		}

		time.Sleep(time.Second * 1)
	}
}

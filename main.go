package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {

	const network = "tcp4"
	const address = "0.0.0.0:12345"

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels, orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(len(proverbs) - 1)
		c.Write([]byte(proverbs[r] + "\n"))
		time.Sleep(time.Second * 3)
	}
}

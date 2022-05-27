package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "Port")
	host = flag.String("h", "localhost", "Host")
)

// connect host:port
// Can Write -> host:port
// Can Read -> host:port
// In the console we send a Hi -> host:port -> Other port can read the "hi"
func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	// We create channel, to have concurrency, and works like a traffic light,
	//just specify when is blocked or when is not.
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{} //to unblock the channel

	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done //Block app until everything is finished
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

}

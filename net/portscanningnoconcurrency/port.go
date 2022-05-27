package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		//Check if a port is open or not
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", i)
		//Too slow
	}
}

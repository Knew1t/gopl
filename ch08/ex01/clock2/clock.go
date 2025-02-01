// Exercise 8.1: Modify clock2 to accept a port number, and write a program,
// clockwall, that acts as a client of several clock servers at once,
// reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices.
// If you have access to geographically distributed computers,
// run instances remotely; otherwise run local instances on different ports with fake time zones.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func getTimeZone(timeZone string) func() time.Time {
	return func() time.Time {
		loc, err := time.LoadLocation(timeZone)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return time.Now().In(loc)
	}
}

var cities = map[string]func() time.Time{
	"8000": getTimeZone("UTC"),
	"5000": getTimeZone("Asia/Tokyo"),
	"3000": getTimeZone("Europe/Dublin"),
	"3400": getTimeZone("Europe/Moscow"),
}

// loc, err := time.LoadLocation("America/New_York")
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("New York Time:", time.Now().In(loc))
// }

func handleConn(c net.Conn) {
	defer c.Close()
	_, portNum, _ := net.SplitHostPort(c.LocalAddr().String())
	for {
		t := cities[portNum]()
		_, err := io.WriteString(c, t.Format("15:04:05\n"))
		if err != nil {
			fmt.Println("Error:", err)
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	portNumber := os.Args[1]
	listener, err := net.Listen("tcp", "localhost:"+portNumber)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

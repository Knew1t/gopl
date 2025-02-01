package main

import (
	"io"
	"log"
	"net"
	// "os"
	"bufio"
	"fmt"
)

// var cities = map[string]func() time.Time{
// 	"8000": getTimeZone("UTC"),
// 	"5000": getTimeZone("Asia/Tokyo"),
// 	"3000": getTimeZone("Europe"),
// 	"3400": getTimeZone("Europe/Moscow"),
// }

var ports = [4]string{"8000", "5000", "3000", "3400"}
var countries = map[string]string{
	"8000": "UTC",
	"5000": "Asia/Tokyo",
	"3000": "Europe",
	"3400": "Europe/Moscow",
}

func Connect(portNumber string) {
	conn, err := net.Dial("tcp", "localhost:"+portNumber)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	// mustCopy(os.Stdout, conn)
	reader := bufio.NewReader(conn)
	for {
		// Read a line from the server
		timeStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Format the output with the country and time
		country := countries[portNumber]
		fmt.Printf("%s (%s): %s", country, portNumber, timeStr)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func main() {

	for _, port := range ports {
		go Connect(port)
	}
	for {
	}
}

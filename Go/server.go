package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":9000")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println(n)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		s := string(buf[0:n])
		v := strings.Split(s, "/")
		for i, z := range v {
			fmt.Print(i)
			fmt.Print(" ")
			fmt.Println(z)
		}
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

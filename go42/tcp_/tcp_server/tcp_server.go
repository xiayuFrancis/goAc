package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	fmt.Println("Server ready to read...")

	for true {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error : ", err)
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(con *net.TCPConn) {
	ipStr := con.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected: " + ipStr)
	}()
	reader := bufio.NewReader(con)
	i := 0

	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))
		time.Sleep(time.Second * 3)

		msg := time.Now().String() + con.RemoteAddr().String() + "Server save hello!\n"
		b := []byte(msg)

		con.Write(b)
		i++

		if i > 10 {
			break
		}
	}
}

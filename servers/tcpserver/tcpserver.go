package tcpserver

import (
	"fmt"
	"net"
	"sync"
)

var (
	tcpListener *net.TCPListener
	done        <-chan struct{}
	port        string
)

func StartUp(quitApplication <-chan struct{}, portNum int) error {
	done = quitApplication
	port = fmt.Sprintf(":%d", portNum)

	addr, err := net.ResolveTCPAddr("tcp", port)

	if err != nil {
		fmt.Printf("error resolving TCP address %v\n", err.Error())
		return err
	}

	tcpListener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("TCP server error on listening: %v\n", err.Error())
		return err
	}

	fmt.Printf("TCP server is listening on %s\n", port)

	//Need to add handler.
	//go tcp_handler.ListenForTCPRequests(done, tcpListener)
	return nil
}

func Shutdown(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Closing TCP connection on port: %s\n", port)
	err := tcpListener.Close()

	if err != nil {
		fmt.Printf("error when shutting down TCP connection with address %v\n", port)
	}
}

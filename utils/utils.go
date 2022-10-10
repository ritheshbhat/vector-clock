package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
)

func HandleRequest(conn net.Conn, logger *govec.GoLog, serverName string) {
	fmt.Printf("before receiving event, log value in %v server is", serverName)
	logger.GetCurrentVC().PrintVC()

	opts := govec.GetDefaultLogOptions()
	var data string
	buf := make([]byte, 1024)
	for {
		_, err := bufio.NewReader(conn).Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		logger.UnpackReceive("Received Message From Client", buf, &data, opts)
		fmt.Printf("vector clock value for %v after receiving request is: ", serverName)
		logger.GetCurrentVC().PrintVC()
		resp := logger.PrepareSend("sending ack to client", 122, opts)
		_, err = conn.Write(resp)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("vector clock value for %v after sending ack is: ", serverName)
		logger.GetCurrentVC().PrintVC()
		defer conn.Close()
	}
}

func SetupConnection(sendingPort, listeningPort string) *net.TCPConn {
	// x := net.DialTimeout("tcp",":"+sendingPort,time.Second * 1)
	rAddr, err := net.ResolveTCPAddr("tcp", ":"+sendingPort)
	ExitIfError(err)
	lAddr, err1 := net.ResolveTCPAddr("tcp", ":"+listeningPort)
	ExitIfError(err)

	conn, errDial := net.DialTCP("tcp", lAddr, rAddr)
	ExitIfError(errDial)
	if (err == nil) && (err1 == nil) && (errDial == nil) {
		return conn
	}
	return nil
}

func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

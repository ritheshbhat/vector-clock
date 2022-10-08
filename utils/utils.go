package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
)

func HandleRequest(conn net.Conn, logger *govec.GoLog) {
	fmt.Println("before receiving event, log value is")
	logger.GetCurrentVC().PrintVC()

	opts := govec.GetDefaultLogOptions()
	var data string
	buf := make([]byte, 1024)
	fmt.Println("handling request...")
	for {
		_, err := bufio.NewReader(conn).Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		logger.UnpackReceive("Received Message From Client", buf, &data, opts)
		fmt.Println("Received ", data)
		fmt.Println("log value post receiving request.")
		logger.GetCurrentVC().PrintVC()
		resp := logger.PrepareSend("sending ack to client", 122, opts)
		conn.Write(resp)
		fmt.Println("sent request to client. log value is.")
		logger.GetCurrentVC().PrintVC()
		break
	}
}

func SetupConnection(sendingPort, listeningPort string) *net.TCPConn {
	rAddr, errR := net.ResolveTCPAddr("tcp", ":"+sendingPort)
	PrintErr(errR)
	lAddr, errL := net.ResolveTCPAddr("tcp", ":"+listeningPort)
	PrintErr(errL)

	conn, errDial := net.DialTCP("tcp", lAddr, rAddr)
	PrintErr(errDial)
	if (errR == nil) && (errL == nil) && (errDial == nil) {
		return conn
	}
	return nil
}

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/DistributedClocks/GoVector/govec"
// 	"github.com/ritheshbhat/vector-clock/firstNode"
// 	"github.com/ritheshbhat/vector-clock/secondNode"
// 	"github.com/ritheshbhat/vector-clock/thirdNode"
// )

// func init() {
// 	 go firstNode.StartFirstNode()
// 	 fmt.Println("starting second node.")
// 	 secondNode.StartSecondNode()
// 	 thirdNode.StartThirdNode()
// 	time.Sleep(time.Millisecond * 200)
// }

// func main() {

// 	fmt.Println("initiating vector clock")

// 	logger := govec.InitGoVector("MyProcess", "LogFile", govec.GetDefaultConfig())

// 	// Encode message, and update vector clock
// 	messagePayload := []byte("sample-payload")
// 	fmt.Println("vector clock value is.")
// 	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())
// 	logger.GetCurrentVC().PrintVC()
// 	// Send message
// 	// connection.Write(vectorClockMessage)

// 	// // In Receiving Process
// 	// connection.Read(vectorClockMessage)

// 	// Decode message, and update local vector clock with received clock
// 	logger.UnpackReceive("Receiving Message", vectorClockMessage, &messagePayload, govec.GetDefaultLogOptions())
// 	logger.GetCurrentVC().PrintVC()
// 	logger.GetCurrentVC().PrintVC()
// 	fmt.Println(logger.GetCurrentVC().FindTicks("1"))

// 	// Log a local event
// 	logger.LogLocalEvent("Example Complete", govec.GetDefaultLogOptions())
// }

package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/DistributedClocks/GoVector/govec"
)

// Hardcoded client/sever port values, and total messages sent
const (
	SERVERPORT = "8080"
	CLIENTPORT = "8081"
	MESSAGES   = 10
)

var done chan int = make(chan int, 2)

func main() {
	go server(SERVERPORT)
	go client(CLIENTPORT, SERVERPORT)
	<-done
	<-done
}

func client(listen, send string) {
	Logger := govec.InitGoVector("client", "clientlogfile", govec.GetDefaultConfig())
	// sending UDP packet to specified address and port
	conn := setupConnection(SERVERPORT, CLIENTPORT)
	opts := govec.GetDefaultLogOptions()

	for i := 0; i < MESSAGES; i++ {
		outgoingMessage := i
		outBuf := Logger.PrepareSend("Sending message to server", outgoingMessage, opts)
		_, errWrite := conn.Write(outBuf)
		printErr(errWrite)

		var inBuf [512]byte
		var incommingMessage int
		n, errRead := conn.Read(inBuf[0:])
		printErr(errRead)
		Logger.UnpackReceive("Received Message from server", inBuf[0:n], &incommingMessage, opts)
		fmt.Printf("GOT BACK : %d\n", incommingMessage)
		time.Sleep(1)
	}
	done <- 1

}

func server(listen string) {

	Logger := govec.InitGoVector("server", "server", govec.GetDefaultConfig())

	fmt.Println("Listening on server....")
	conn, err := net.ListenPacket("udp", ":"+listen)
	printErr(err)

	var buf [512]byte

	var n, nMinOne, nMinTwo int
	n = 1
	nMinTwo = 1
	nMinTwo = 1
	opts := govec.GetDefaultLogOptions()

	for i := 0; i < MESSAGES; i++ {
		_, addr, err := conn.ReadFrom(buf[0:])
		var incommingMessage int
		Logger.UnpackReceive("Received Message From Client", buf[0:], &incommingMessage, opts)
		fmt.Printf("Received %d\n", incommingMessage)
		printErr(err)

		switch incommingMessage {
		case 0:
			nMinTwo = 0
			n = 0
			break
		case 1:
			nMinOne = 0
			n = 1
			break
		default:
			nMinTwo = nMinOne
			nMinOne = n
			n = nMinOne + nMinTwo
			break
		}

		outBuf := Logger.PrepareSend("Replying to client", n, opts)

		conn.WriteTo(outBuf, addr)
		time.Sleep(1)
	}
	conn.Close()
	done <- 1

}

func setupConnection(sendingPort, listeningPort string) *net.UDPConn {
	rAddr, errR := net.ResolveUDPAddr("udp4", ":"+sendingPort)
	printErr(errR)
	lAddr, errL := net.ResolveUDPAddr("udp4", ":"+listeningPort)
	printErr(errL)

	conn, errDial := net.DialUDP("udp", lAddr, rAddr)
	printErr(errDial)
	if (errR == nil) && (errL == nil) && (errDial == nil) {
		return conn
	}
	return nil
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

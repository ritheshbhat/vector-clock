package main

import (
	"fmt"
	"net"
	"time"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/ritheshbhat/vector-clock/firstNode"
	"github.com/ritheshbhat/vector-clock/secondNode"
	"github.com/ritheshbhat/vector-clock/thirdNode"
	"github.com/ritheshbhat/vector-clock/utils"
)

// Hardcoded client/sever port values
const (
	FirstNodeServerPort  = "8080"
	SecondNodeServerPort = "8081"
	ThirdNodeServerPort  = "8082"
	FirstNodeClientPort  = "8083"
	SecondNodeClientPort = "8084"
	ThirdNodeClientPort  = "8085"
	UniversalClientPort  = "8086"
)

func init() {
	go firstNode.StartServer(FirstNodeServerPort)
	go secondNode.StartServer(SecondNodeServerPort)
	go thirdNode.StartServer(ThirdNodeServerPort)

	time.Sleep(time.Second * 2)
	fmt.Println("3 servers and up and listening for incoming requets")
}

var conn1 *net.TCPConn
var conn2 *net.TCPConn
var conn3 *net.TCPConn

type clientConfigs struct {
	name  string
	log   *govec.GoLog
	opts  govec.GoLogOptions
	conn1 *net.TCPConn
	conn2 *net.TCPConn
	conn3 *net.TCPConn
}

func loadClientLogConfigs(name string) (*govec.GoLog, govec.GoLogOptions) {
	logger := govec.InitGoVector(name, "clientLogFile", govec.GetDefaultConfig())
	opts := govec.GetDefaultLogOptions()
	return logger, opts

}

//	func getTCPConnection(userInput string, name string) (string, string) {
//		if userInput == "1" {
//			return client(FirstNodeClientPort,FirstNodeServerPort, name)
//			return FirstNodeClientPort, FirstNodeServerPort
//		} else if userInput == "2" {
//			return SecondNodeClientPort, SecondNodeServerPort
//		} else if userInput == "3" {
//			return ThirdNodeClientPort, ThirdNodeServerPort
//		}
//		return "0", "0"
//	}

func main() {
	conn1 = client(FirstNodeClientPort, FirstNodeServerPort, "client")
	conn2 = client(SecondNodeClientPort, SecondNodeServerPort, "client")
	conn3 = client(ThirdNodeClientPort, ThirdNodeServerPort, "client")
	defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()

	logger, opts := loadClientLogConfigs("client")
	uc := clientConfigs{
		"client",
		logger,
		opts,
		conn1,
		conn2,
		conn3,
	}

	//TODO: decouple logger for each client, so their event is stored, otherwise everytime a client func is called, the logger is reset.
	for {
		time.Sleep(time.Second * 3)
		fmt.Println("Server id's available to send messages are 1,2,3: ")
		var in string
		fmt.Scanln(&in)
		if in == "1" {
			go makeServerCall(uc.conn1, uc.log, uc.name, uc.opts)
		} else if in == "2" {
			go makeServerCall(uc.conn2, uc.log, uc.name, uc.opts)
		} else if in == "3" {
			go makeServerCall(uc.conn3, uc.log, uc.name, uc.opts)

		}

	}

}

func client(listen, send string, clientName string) *net.TCPConn {
	conn := utils.SetupConnection(send, listen)
	return conn
}

func makeServerCall(conn *net.TCPConn, Logger *govec.GoLog, clientName string, opts govec.GoLogOptions) {
	fmt.Printf("current vector value before sending request for  %v is: ", clientName)
	Logger.GetCurrentVC().PrintVC()
	outgoingMessage := clientName
	outBuf := Logger.PrepareSend("Sending message to server", outgoingMessage, opts)
	_, errWrite := conn.Write(outBuf)
	if errWrite != nil {
		utils.PrintErr(errWrite)
	}
	fmt.Printf("vector clock value post sending request by %v", clientName)
	Logger.GetCurrentVC().PrintVC()
	var inBuf [512]byte
	var incommingMessage int
	n, errRead := conn.Read(inBuf[0:])
	utils.PrintErr(errRead)
	Logger.UnpackReceive("Received Message from server", inBuf[0:n], &incommingMessage, opts)
	fmt.Printf("vector clock value after receiving request from server by. %v ", clientName)
	Logger.GetCurrentVC().PrintVC()
}

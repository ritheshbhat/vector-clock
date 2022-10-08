package main

import (
	"fmt"
	"time"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/ritheshbhat/vector-clock/firstNode"
	"github.com/ritheshbhat/vector-clock/secondNode"
	"github.com/ritheshbhat/vector-clock/thirdNode"
	"github.com/ritheshbhat/vector-clock/utils"
)

// func init() {
// 	 go firstNode.StartFirstNode()
// 	 fmt.Println("starting second node.")
// 	 secondNode.StartSecondNode()
// 	 thirdNode.StartThirdNode()
// 	time.Sleep(time.Millisecond * 200)
// }

func mainz() {

	fmt.Println("initiating vector clock")

	logger := govec.InitGoVector("MyProcess", "LogFile", govec.GetDefaultConfig())

	// Encode message, and update vector clock
	messagePayload := []byte("sample-payload")
	fmt.Println("vector clock value is.")
	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())
	logger.GetCurrentVC().PrintVC()
	// Send message
	// connection.Write(vectorClockMessage)

	// // In Receiving Process
	// connection.Read(vectorClockMessage)

	// Decode message, and update local vector clock with received clock
	logger.UnpackReceive("Receiving Message", vectorClockMessage, &messagePayload, govec.GetDefaultLogOptions())
	logger.GetCurrentVC().PrintVC()
	logger.GetCurrentVC().PrintVC()
	fmt.Println(logger.GetCurrentVC().FindTicks("1"))

	// Log a local event
	logger.LogLocalEvent("Example Complete", govec.GetDefaultLogOptions())
}

// Hardcoded client/sever port values, and total messages sent
const (
	FirstNodeServerPort  = "8080"
	SecondNodeServerPort = "8081"
	ThirdNodeServerPort  = "8082"
	FirstNodeClientPort  = "8083"
	SecondNodeClientPort = "8084"
	ThirdNodeClientPort  = "8085"

	MESSAGES = 10
)

var done chan int = make(chan int, 2)

func main() {
	go firstNode.StartServer(FirstNodeServerPort, done)
	go secondNode.StartServer(SecondNodeServerPort, done)
	go thirdNode.StartServer(ThirdNodeServerPort, done)

	time.Sleep(time.Second * 5)
	//TODO: understand way to receive obj from async func and then interact with it.
	go client(FirstNodeClientPort, FirstNodeServerPort, "firstNodeClient")
	go client(SecondNodeClientPort, SecondNodeServerPort, "secondNodeClient")
	go client(ThirdNodeClientPort, ThirdNodeServerPort, "thirdNodeClient")
	time.Sleep(time.Second * 10)

	<-done
	<-done
	<-done
}

func client(listen, send string, clientName string) {
	Logger := govec.InitGoVector(clientName, "clientlogfile", govec.GetDefaultConfig())
	conn := utils.SetupConnection(send, listen)
	opts := govec.GetDefaultLogOptions()

	fmt.Printf("initiating msg send.. vector clock value for %v", clientName)
	Logger.GetCurrentVC().PrintVC()
	outgoingMessage := clientName
	outBuf := Logger.PrepareSend("Sending message to server", outgoingMessage, opts)
	_, errWrite := conn.Write(outBuf)
	utils.PrintErr(errWrite)
	fmt.Printf("vector clock value post sending request by %v", clientName)
	Logger.GetCurrentVC().PrintVC()
	var inBuf [512]byte
	var incommingMessage int
	n, errRead := conn.Read(inBuf[0:])
	utils.PrintErr(errRead)
	Logger.UnpackReceive("Received Message from server", inBuf[0:n], &incommingMessage, opts)
	fmt.Printf("GOT BACK : %d\n", incommingMessage)
	fmt.Printf("vector clock value after receiving request from server by. %v ", clientName)
	Logger.GetCurrentVC().PrintVC()

	time.Sleep(1)
	done <- 1

}

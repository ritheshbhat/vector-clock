package main

import (
	"fmt"

	"github.com/DistributedClocks/GoVector/govec"
)

func init(){
	fmt.Println("initiating client and server.")
	go firstNode()
	go secondNode()
	go thirdNode()
}

func main() {
	
	fmt.Println("initiating vector clock")

	logger := govec.InitGoVector("MyProcess", "LogFile", govec.GetDefaultConfig())

	// Encode message, and update vector clock
	messagePayload := []byte("sample-payload")
	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())

	// Send message
	connection.Write(vectorClockMessage)

	// In Receiving Process
	connection.Read(vectorClockMessage)

	// Decode message, and update local vector clock with received clock
	logger.UnpackReceive("Receiving Message", vectorClockMessage, &messagePayload, govec.GetDefaultLogOptions())

	// Log a local event
	logger.LogLocalEvent("Example Complete", govec.GetDefaultLogOptions())
}

package thirdNode

import (
	"fmt"

	"github.com/DistributedClocks/GoVector/govec"
)

func StartThirdNode() {
	logger := govec.InitGoVector("thirdNode", "LogFile", govec.GetDefaultConfig())
	// Encode message, and update vector clock
	messagePayload := []byte("sample-payload")
	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())
	fmt.Println(string(vectorClockMessage))

}

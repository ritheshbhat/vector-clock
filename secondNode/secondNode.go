package secondNode

import (
	"fmt"

	"github.com/DistributedClocks/GoVector/govec"
)

func StartSecondNode() {
	logger := govec.InitGoVector("secondNode", "LogFile", govec.GetDefaultConfig())

	// Encode message, and update vector clock
	messagePayload := []byte("sample-payload")
	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())
	fmt.Println(string(vectorClockMessage))

}

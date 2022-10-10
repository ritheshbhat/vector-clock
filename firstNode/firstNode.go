package firstNode

import (
	"fmt"
	"net"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/ritheshbhat/vector-clock/utils"
)

func StartServer(listen string) {
	serverName := "firstNode"
	logger := govec.InitGoVector(serverName, "server", govec.GetDefaultConfig())
	fmt.Printf("first Node server vector clock event is: ")
	logger.GetCurrentVC().PrintVC()
	conn, err := net.Listen("tcp", ":"+listen)
	if err!=nil{
		fmt.Println("err is", err)
	}
	defer conn.Close()

	for {
		// Listen for an incoming connection.
		x, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		 go utils.HandleRequest(x, logger, serverName)
	}

}

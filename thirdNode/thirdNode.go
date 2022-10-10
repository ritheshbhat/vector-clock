package thirdNode

import (
	"fmt"
	"net"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/ritheshbhat/vector-clock/utils"
)

func StartServer(listen string) {
	name := "thirdNode"

	logger := govec.InitGoVector(name, "server", govec.GetDefaultConfig())
	fmt.Println("third node vector clock event is: ")
	logger.GetCurrentVC().PrintVC()
	conn, _ := net.Listen("tcp", ":"+listen)
	defer conn.Close()
	for {
		// Listen for an incoming connection.
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go utils.HandleRequest(conn, logger,name)
	}


}

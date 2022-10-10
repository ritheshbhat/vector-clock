package secondNode

import (
	"fmt"
	"net"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/ritheshbhat/vector-clock/utils"
)

func StartServer(listen string) {
	name := "secondNode"

	logger := govec.InitGoVector(name, "server", govec.GetDefaultConfig())
	fmt.Println("second node server vector clock event is: ")
	logger.GetCurrentVC().PrintVC()
	conn, err := net.Listen("tcp", ":"+listen)
	if err!=nil{
		fmt.Println("err is", err)
	}
	defer conn.Close()
	for {
		// Listen for an incoming connection.
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go utils.HandleRequest(conn, logger,"secondNode")
	}

}

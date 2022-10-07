package firstNode

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DistributedClocks/GoVector/govec"
	"github.com/gorilla/mux"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func firstNode(w http.ResponseWriter, r *http.Request) {
	logger := govec.InitGoVector("firstNode", "LogFile", govec.GetDefaultConfig())
	// Encode message, and update vector clock
	messagePayload := []byte("sample-payload")
	vectorClockMessage := logger.PrepareSend("Sending Message", messagePayload, govec.GetDefaultLogOptions())
	fmt.Println(string(vectorClockMessage))

}
func StartFirstNode() {

	router := mux.NewRouter()
	router.HandleFunc("/first-node", firstNode)
	s := &server{}
	http.Handle("/", s)
	fmt.Println("serving node 1.")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

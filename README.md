# vector-clock
 replicate vector clock implementation in N node Distributed system

Install Go: `https://go.dev/dl/`


Configure GOPATH

`export GOPATH=~/<path-where-binaries-are-installed-for-go>`

### Example to create GOPATH post go installation:
`mkdir go`

`cd go`

`mkdir src pkg bin`

`cd src; <unzip the uploaded folder>`

#export GOPATH to `go` folder location.

`go mod tidy`


# Start simulated 3 distributed nodes
`go run main.go`

# Start interacting with the application, press 1 / 2 / 3 to communicate accordingly with the servers 1 / 2 / 3 respectively.

# Functionalities:
`1. establish TCP connection with 3 servers listening on 3 different ports`

`2. Display vector clock event for each server`

`3. Display vector clock events of all the nodes that have been communicated with`


### The code is pushed to a private repository in github.
https://github.com/ritheshbhat/vector-clock
Do let me know if you'd need access to this repo to clone.

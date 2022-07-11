package main

import board "github.com/go-funcards/board-service/cmd"

//go:generate protoc -I proto --go_out=./proto/v1 --go-grpc_out=./proto/v1 proto/v1/board.proto

func main() {
	board.Execute()
}

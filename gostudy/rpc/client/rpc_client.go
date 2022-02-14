package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// var code string = "tcp"
var code string = "http"

func main() {
	if code == "tcp" {
		client, err := rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing", err)
		}
		var reply string

		err = client.Call("Service.Hello", "TCP", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}

	if code == "http" {
		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing", err)
		}
		var reply string

		err = client.Call("Service.Hello", "HTTP", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}

}

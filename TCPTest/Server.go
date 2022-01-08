package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	dstream, err := net.Listen("tcp", ":8080")

	if err != nil{
		fmt.Println(err)
		return
	}

	defer dstream.Close()

	for{
		connection, err := dstream.Accept()
		if err != nil{
			fmt.Println(err)
			return
		}

		//goroutine which is like coroutine in c#
		go Work(connection)
	}
}

func Work(connection net.Conn){
	for{
		data, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
	connection.Close()
}
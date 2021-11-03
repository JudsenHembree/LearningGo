
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import ("fmt"
		"os")

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

var users map[string]string

func main() {
	//fmt.Println(len(os.Args))
	var user string = os.Args[1]
	var password string= os.Args[2]
	users = make(map[string]string)
	users["Judsen"]="8888"
	users["Seven"]="7777"

	if(len(os.Args)!=3){
		fmt.Println("Usage: [username] [password]")
		return
	}
	if _, ok := users[user];ok{
		if(password==users[user]){
			fmt.Println("Access Granted to "+ user)
			return
		}
		fmt.Println("Password incorrect for "+ user)
		return
	}
	fmt.Println("Access denied for "+ user)
}
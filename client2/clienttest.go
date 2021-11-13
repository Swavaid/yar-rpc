package main

import (
	"dis_sys"
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8888")
	if err!=nil{
		fmt.Println("接收error")
	}

	cli := dis_sys.NewClient(conn)
	// 声明函数原型
	var query func (str string) string

	cli.CallRPC("uppercase", &query)
	n:=query("hellorpc")
	fmt.Println(n)
}
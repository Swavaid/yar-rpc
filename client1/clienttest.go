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
	var query func (a float64,b float64) float64

	cli.CallRPC("sum", &query)
	n:=query(8.89,9.0)
	fmt.Println(n)
}

package main

import (
	"dis_sys"
	"encoding/gob"
	"fmt"
	"strings"
)

type Start struct {

}

func sum(a float64,b float64) float64 {
	fmt.Printf("the sum of a and b=%v",a+b)
	return a+b
}

func uppercase(str string) string {
	strUp:=strings.ToUpper(str)
	return strUp
}


func main(){
	gob.Register(Start{})
	server1:= dis_sys.NewServer("localhost:8888")
	server1.Register("sum",sum)
	server1.Register("uppercase",uppercase)
	server1.Run()
}
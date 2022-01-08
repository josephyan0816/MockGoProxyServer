package main

import (
	"fmt"
	"mockgoproxy/handler"
)

func main()  {
	myproxy:=handler.NewProxyServer()
	fmt.Println(myproxy)
}

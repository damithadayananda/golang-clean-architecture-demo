package main

import (
	"fmt"
	"golang-clean-architecture/interface/http/inbound"
)

func main()  {
	fmt.Println("welcome to clean arch demo")
	inbound.InitServer()
}

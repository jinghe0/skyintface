package main

import "fmt"
import "network"

func main() {
     fmt.Println("Server Start")
     filter := new(network.NetFilter)
     filter.Init()
     filter.Run()  //Todo 不知道为什么用 go后，这里
     fmt.Println("Server Start 2")
}

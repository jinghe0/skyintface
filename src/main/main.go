package main

import (
 "fmt"
 "network"
 "storage"
)

func main() {
     fmt.Println("Server Start")

     redisHelper := new(storage.RedisHelper)
     redisHelper.Init()

     filter := new(network.NetFilter)
     filter.Init()
     filter.Run()  
         
     fmt.Println("Server Start 2")
}

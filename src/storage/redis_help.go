package storage

import (
    "fmt"
    "github.com/alphazero/Go-Redis"
)

type RedisHelper struct {
    client  redis.Client
} 

func (helper *RedisHelper) Init() {
    spec := redis.DefaultSpec()
    client, err := redis.NewSynchClientWithSpec(spec)
    if err != nil {
        fmt.Println("error on connect redis server")
        return
    }

    helper.client = client
}

// func (helper *RedisHelper) Get(key string) (value string, err Error){
//     result, err := helper.client.Get(dbkey)
//     if err != nil {
//         fmt.Println("Get Key fail")
//         return nil, err
//     }

//     value = string(result)
//     return value, err
// }



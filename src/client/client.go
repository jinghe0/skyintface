package main

import (
    "fmt"
    "os"
    "bytes"
    "net"
    "io"
    "encoding/binary"
    "game"
    "github.com/golang/protobuf/proto"
)

func main() {
//  if len(os.Args) != 2 {
//      fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
//      os.Exit(1)
//  }

    addr := "127.0.0.1:9090"
    
    conn, err := net.Dial("tcp", addr)
    checkError(err)

    // _,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    // checkError(err)

    //data := "Hello world"
    loginReq := &game.CSLoginReq {
        OpenId : proto.String("denny"),     
    }

    data, err := proto.Marshal(loginReq)
    if err != nil {
        fmt.Println("marshaling error : ", err)
    }

    msg := make([]byte, 4+len(data))

    binary.BigEndian.PutUint16(msg, uint16(2+len(data))) 
    binary.BigEndian.PutUint16(msg[2:], uint16(501))
    copy(msg[4:], data)

     _,err = conn.Write(msg)
    checkError(err)

    result, err := readFully(conn)
    checkError(err)

    msg2 := make([]byte, 4+len(data))

    binary.BigEndian.PutUint16(msg2, uint16(2+len(data))) 
    binary.BigEndian.PutUint16(msg2[2:], uint16(501))
    copy(msg2[4:], data)

     _,err = conn.Write(msg2)
    checkError(err)

    
    fmt.Println(string(result))
    
    os.Exit(0)
}

func checkError(err error){
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
    
func readFully(conn net.Conn)([]byte, error) {
    defer conn.Close()
    
    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        fmt.Println("Read socket")
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
        }
        fmt.Println("Server close")
        return nil, err
    }
    
    return result.Bytes(), nil
}
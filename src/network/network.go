package network

import (
    "fmt"
    "net"
    "os"
)

type Filter interface {
    Init()
    ReadMsg() ([]byte, error)
    WriteMsg(args ...[]byte) error
}

type NetFilter struct {
    conn net.Conn
    ln net.Listener
}

func (filter *NetFilter) Write(b []byte){
    
}

func (filter *NetFilter) Read(b []byte){

}

func (filter *NetFilter) Init(){
    fmt.Println("Network Filter Init")
    ln, err := net.Listen("tcp", ":9090")
    filter.ln = ln
    if err != nil {
        fmt.Println("Listen error")
        os.Exit(-1)
    }    
}

func handleConn(conn net.Conn){
    msgParser := new(MsgParser)
    tcpConn := newTcpConn(conn, msgParser)
    tcpConn.ReadMsg()
}

func (filter *NetFilter) Run(){
    fmt.Println("Accept Loop Begin")
    for {
        conn, err := filter.ln.Accept()
        fmt.Println("Socket Accept")
        if err != nil {
            fmt.Println("Accept error")
            os.Exit(-1)
        }
        go handleConn(conn)  //利用携程来做接收处理
    }
    fmt.Println("Accept Loop End")
}














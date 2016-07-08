package network

import (
    "fmt"
    "net"
    "encoding/binary"
    "game"
    "github.com/golang/protobuf/proto"
)

type TcpConn struct {
    conn net.Conn
    msgParser *MsgParser
    marshal *Marshal
}

func newTcpConn(conn net.Conn, msgParser *MsgParser) *TcpConn{
    tcpConn := new(TcpConn)
    tcpConn.conn = conn
    tcpConn.msgParser = msgParser

    return tcpConn
}

func (tcpConn *TcpConn) Read(b []byte) (int, error) {
    return tcpConn.conn.Read(b)
}

func Process(msgType uint32, data []byte){
    fmt.Println("Process Data  Type: ", msgType)

   loginReq := &game.CSLoginReq{}
   err := proto.Unmarshal(data, loginReq)
   if err != nil {
        fmt.Println("unmarshaling error : ", err)
   }

   fmt.Println("login openid: ", loginReq.GetOpenId())

    //fmt.Println("Data : ", string(buf))
}

func (tcpConn *TcpConn) ReadMsg()  {
    fmt.Println("Read Msg ---------------- ")
    //data := tcpConn.msgParser.Read(tcpConn.conn)
    //data := tcpConn.msgParser.Read(tcpConn.conn)
    msgParser := tcpConn.msgParser
    data, _ := msgParser.Read(tcpConn.conn)

    fmt.Println("After recv full data")
    msgType := uint32(binary.BigEndian.Uint16(data[:2]))

    Process(msgType, data[2:])
    //marshal.unpack(data)
    //解析报文，并根据消息号，做相应逻辑跳转
}
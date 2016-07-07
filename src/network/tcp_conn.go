package network

import (
    "fmt"
    "net"
    "encoding/binary"
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

func Process(msgType uint32, buf []byte){
    fmt.Println("Process Data %d", msgType)
    fmt.Println("Data : ", string(buf))
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
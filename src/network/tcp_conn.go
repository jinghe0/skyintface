package network

import (
    "fmt"
    "os"
    "net"
    "encoding/binary"
    "game"
    "github.com/golang/protobuf/proto"
    "logic"
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

func (tcpConn *TcpConn) Process(msgType uint32, data []byte){
    fmt.Println("Process Data  Type: ", msgType)

    loginReq := &game.CSLoginReq{}
    err := proto.Unmarshal(data, loginReq)
    if err != nil {
        fmt.Println("unmarshaling error : ", err)
    }

    agentInstance := new(logic.Agent)
    agentInstance.Init(loginReq.GetOpenId())

    fmt.Println("login openid: ", loginReq.GetOpenId())

    tcpConn.WriteMsg()


    //fmt.Println("Data : ", string(buf))
}

func (tcpConn *TcpConn) ReadMsg()  {
    fmt.Println("Read Msg")
    //data := tcpConn.msgParser.Read(tcpConn.conn)
    //data := tcpConn.msgParser.Read(tcpConn.conn)
    msgParser := tcpConn.msgParser
    data, _ := msgParser.Read(tcpConn.conn)

    fmt.Println("After recv full data")
    msgType := uint32(binary.BigEndian.Uint16(data[:2]))

    tcpConn.Process(msgType, data[2:])


    //marshal.unpack(data)
    //解析报文，并根据消息号，做相应逻辑跳转
}

func (tcpConn *TcpConn) WriteMsg() {
    fmt.Println("Write Msg")

    loginReq := &game.CSLoginReq {
        OpenId : proto.String("denny"),     
    }

    data, err := proto.Marshal(loginReq)

    msg := make([]byte, 4+len(data))

    binary.BigEndian.PutUint16(msg, uint16(2+len(data))) 
    binary.BigEndian.PutUint16(msg[2:], uint16(501))
    copy(msg[4:], data)

     _,err = tcpConn.conn.Write(msg)
    //checkError(err)
     if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
     }
}
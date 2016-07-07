package network

import (
    "net"
    "fmt"
    "io"
    "encoding/binary"
)

const headLen = 2

type Parser interface {
    Read(conn net.Conn)()
    Write(conn net.Conn)()
}

type Busi interface {
    DoReq([] byte)
    DoResp([] byte)
}

type MsgParser struct {
    //通过序列化后解析成相应结构抛给业务层 (json, protobuf, flatbuffer)
}

func (parser* MsgParser) Read(conn net.Conn) ([]byte, error){
    var b [2]byte
    bufLen := b[:2]
    _, err := io.ReadFull(conn, bufLen)
    if err != nil {
        return nil, err
    }

    msgLen := uint32(binary.BigEndian.Uint16(bufLen))

    fmt.Println("Msg Len : ", msgLen)

    msgData := make([]byte, msgLen)
    if _, err := io.ReadFull(conn, msgData); err != nil {
        return nil, err
    }
    
    return msgData, nil
}

func (parser* MsgParser) Write(conn net.Conn){
    fmt.Println("Write data")
}








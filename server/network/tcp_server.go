package network

import (
    "server/log"
	"net"
	"sync"
	"time"
)

type TCPServer struct {
    Addr string
    MaxConnNum uint
    PendingWriteNum uint
    NewAgent func(*TCPConn) Agent
    ln net.Listener
    conns ConnSet
    mutexConns sync.Mutex
    wgLn sync.WaitGroup
    wgConns sync.WaitGroup

    LenMsgLen uint
    MinMsgLen uint32
    MaxMsgLen uint32
    LittleEndian bool
    msgParser *MsgParser
}

func (server *TCPServer) Start() {
    server.init()
    go server.run()
}

func (server *TCPServer) init() {
    ln, err := net.Listen("tcp", server.Addr)

    if err != nil {
        log.Fatal("%v", err)
    }

    	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		log.Release("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}

	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
		log.Release("invalid PendingWriteNum, reset to %v", server.PendingWriteNum)
	}

	if server.NewAgent == nil {
		log.Fatal("NewAgent must not be nil")
	}

    server.ln = ln

    server.conns = make(ConnSet)

    msgParser := NewMsgParser()
	msgParser.SetMsgLen(server.LenMsgLen, server.MinMsgLen, server.MaxMsgLen)
	msgParser.SetByteOrder(server.LittleEndian)
	server.msgParser = msgParser
}

func (server *TCPServer) run() {
    server.wgLn.Add(1)

    defer server.wgLn.Done()

    var tempDelay time.Duration

    for {
        conn, err := server.ln.Accept()

        if err != nil {
            if ne, ok := err.(net.Error); ok && ne.Temporary() {
                if tempDelay == 0 {
                    tempDelay = 5 * time.Millisecond
                } else {
                    tempDelay *= 2
                }
                if max := 1 * time.Second; tempDelay > max {
                    tempDelay = max
                }
            }
        }
    }
}

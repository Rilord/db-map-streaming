package network

import (
	"game-server/log"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	sync.Mutex
	Addr            string
	ConnNum         uint
	ConnectInterval time.Duration
	PendingWriteNum uint
	AutoReconnect   bool
	NewAgent        func(*TCPConn) Agent
	conns           ConnSet
	wg              sync.WaitGroup
	closeFlag       bool

	LenMsgLen    uint
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	msgParser    *MsgParser
}

package network

import (
	"game-server/log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	sync.Mutex
	Addr             string
	ConnNum          uint
	ConnectInterval  time.Duration
	PendingWriteNum  uint
	MaxMsgLen        uint32
	HandshakeTimeout time.Duration
	AutoReconnect    bool
	NewAgent         func(*WSConn) Agent
	dialer           websocket.Dialer
	conns            WebsocketConnSet
	wg               sync.WaitGroup
	closeFlag        bool
}

func (client *WSClient) Start() {
	client.init()

}

func (client *WSClient) init() {
	client.Lock()

	defer client.Unlock()

	if client.ConnNum <= 0 {
		client.ConnNum = 1
		log.Release("invalid ConnNum, setting to %v", client.ConnNum)
	}
	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		log.Release("invalid ConnectInterval, setting to %v", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		log.Release("invalid PendingWriteNum, setting to %v", client.PendingWriteNum)
	}
	if client.MaxMsgLen <= 0 {
		client.MaxMsgLen = 4096
		log.Release("invalid MaxMsgLen, setting to %v", client.MaxMsgLen)
	}
	if client.HandshakeTimeout <= 0 {
		client.HandshakeTimeout = 10 * time.Second
		log.Release("invalid HandshakeTimeout, setting to %v", client.HandshakeTimeout)
	}
	if client.NewAgent == nil {
		log.Fatal("NewAgent must not be nil")
	}
	if client.conns != nil {
		log.Fatal("client is running")
	}

	client.conns = make(WebsocketConnSet)

	client.closeFlag = false
	client.dialer = websocket.Dialer{
		HandshakeTimeout: client.HandshakeTimeout,
	}
}

func (client *WSClient) dial() *websocket.Conn {
	for {
		conn, _, err := client.dialer.Dial(client.Addr, nil)
		if err == nil || client.closeFlag {
			return conn
		}

		log.Release("connect to %v error: %v", client.Addr, err)
		time.Sleep(client.ConnectInterval)
		conn
	}
}

func (client *WSClient) connect() {
    defer client.wg.Done()

reconnect:
    conn := client.dial()
    if conn == nil {
        return
    }




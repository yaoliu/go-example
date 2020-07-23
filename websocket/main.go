//__author__ = "YaoYao"
//Date: 2020/6/29
package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"net"
	"sync"
	"time"
)

type Server interface {
	ListenAndServe() error
	Serve(listener net.Listener) error
	Scheme() string
	Close() error
}

type Handler interface {
	Handle(in net.Conn)
}

type EchoHandler struct{}

func (e EchoHandler) Handle(in net.Conn) {
	b := make([]byte, 100)
	for {
		n, err := in.Read(b)
		if err != nil {
			fmt.Println(err)
		}

		n, err = in.Write(b[:n])
		if err != nil {
			fmt.Println(err)
		}
	}
}

var _ Handler = &EchoHandler{}

type WSServer struct {
	Addr         string `json:"addr"`
	Path         string `json:"path"`
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	mu           sync.Mutex
	handler      Handler
	listener     net.Listener
	conns        map[net.Conn]bool
	MaxConns     int
}
type wsconn struct {
	c net.Conn
}

func (w wsconn) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (w wsconn) Write(b []byte) (n int, err error) {
	panic("implement me")
}

func (w wsconn) Close() error {
	return w.c.Close()
}

func (w wsconn) LocalAddr() net.Addr {
	return w.c.LocalAddr()
}

func (w wsconn) RemoteAddr() net.Addr {
	return w.c.RemoteAddr()
}

func (w wsconn) SetDeadline(t time.Time) error {
	return w.c.SetDeadline(t)
}

func (w wsconn) SetReadDeadline(t time.Time) error {
	return w.c.SetReadDeadline(t)
}

func (w wsconn) SetWriteDeadline(t time.Time) error {
	return w.c.SetWriteDeadline(t)
}

var _ Server = &WSServer{}

func NewWSServer() Server {
	return &WSServer{}
}

func (wss *WSServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", "")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := listener.Close()
		fmt.Println(err)
	}()
	return wss.Serve(listener)
}

func (wss *WSServer) Scheme() string {
	return "websocket"
}

func (wss *WSServer) Serve(l net.Listener) error {
	wss.mu.Lock()
	wss.listener = l
	wss.mu.Unlock()
	upgrader := ws.Upgrader{}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func() {
			_, err := upgrader.Upgrade(conn)
			if err != nil {
				fmt.Println(err)
			}
			c := &wsconn{c: conn}

			defer func() {
				wss.mu.Lock()
				delete(wss.conns, c)
				wss.mu.Unlock()
			}()
			wss.handler.Handle(c)
		}()
	}
}

func (wss *WSServer) Close() error {
	wss.closeListener()
	return wss.closeAllConns()
}

func (wss *WSServer) closeListener() {
	wss.mu.Lock()
	wss.listener.Close()
	wss.listener = nil
	wss.mu.Unlock()
}

func (wss *WSServer) closeAllConns() error {
	for conn := range wss.conns {
		conn.Close()
	}
	return nil
}

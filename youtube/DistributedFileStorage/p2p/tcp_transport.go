package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
  listenAddress string
  listener      net.Listener

  mu            sync.RWMutex // mu là một biến đồng bộ, dùng để đồng bộ hóa việc truy cập đến các biến
  peers         map[net.Addr] Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
  return &TCPTransport{
    listenAddress: listenAddr,
  }
}

func (t *TCPTransport) ListenAndAccept() error {
  var err error
  t.listener, err = net.Listen("tcp", t.listenAddress)
  if err != nil {
    return err
  }

  go t.startAcceptLoop()

  return nil
}

func (t *TCPTransport) startAcceptLoop() {
  for {
    conn, err := t.listener.Accept()
    if err != nil {
      fmt.Printf("[-] Error accepting connection: %s\n", err)
    }

    go t.handleConn(conn)
  }
}

func (t *TCPTransport) handleConn(conn net.Conn) {
  // sao phải dung %+v thay vì %v? %v không đủ để in ra tất cả các thông tin của một biến
  fmt.Printf("[+] New connection %+v\n", conn) 

} 


package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer là một peer sử dụng TCP để giao tiếp với các peer khác
type TCPPeer struct {
  conn      net.Conn
  // if we dial and retrieve a connection => outbound = true (Nếu chủ động tạo kết nối đến peer khác bằng Dial())
  // if we accept and retrieve a connection => outbound = false (Nếu nhận được kết nối từ peer khác bằng Accept())
  outbound  bool // xác định xem peer này có phải là peer gửi request không
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
  return &TCPPeer{
    conn:     conn,
    outbound: outbound,
  }
}

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
  peer := NewTCPPeer(conn, true)

  // sao phải dung %+v thay vì %v? %v không đủ để in ra tất cả các thông tin của một biến
  fmt.Printf("[+] New connection %+v\n", peer) 

} 


package rate

import (
	"io"
	"net"
)

type rateConn struct {
	conn net.Conn
	rate *Rate
}

func NewRateConn(conn net.Conn, rate *Rate) io.ReadWriteCloser {
	return &rateConn{
		conn: conn,
		rate: rate,
	}
}

func (s *rateConn) Read(b []byte) (n int, err error) {
	n, err = s.conn.Read(b)
	if s.rate != nil {
		s.rate.Get(int64(n))
	}
	return
}

func (s *rateConn) Write(b []byte) (n int, err error) {
	n, err = s.conn.Write(b)
	if s.rate != nil {
		s.rate.Get(int64(n))
	}
	return
}

func (s *rateConn) Close() error {
	return s.conn.Close()
}

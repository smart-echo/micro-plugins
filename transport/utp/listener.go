package utp

import (
	"bufio"
	"encoding/gob"
	"net"
	"time"

	"github.com/smart-echo/micro/logger"
	"github.com/smart-echo/micro/transport"
)

func (u *utpListener) Addr() string {
	return u.l.Addr().String()
}

func (u *utpListener) Close() error {
	return u.l.Close()
}

func (u *utpListener) Accept(fn func(transport.Socket)) error {
	var tempDelay time.Duration

	for {
		c, err := u.l.Accept()
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
				logger.Errorf("utp: Accept error: %v; retrying in %v\n", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}

		encBuf := bufio.NewWriter(c)

		sock := &utpSocket{
			timeout: u.t,
			conn:    c,
			encBuf:  encBuf,
			enc:     gob.NewEncoder(encBuf),
			dec:     gob.NewDecoder(c),
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					sock.Close()
				}
			}()

			fn(sock)
		}()
	}
}

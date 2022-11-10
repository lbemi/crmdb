package asset

import (
	"bufio"
	"github.com/gorilla/websocket"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"golang.org/x/crypto/ssh"
	"time"
	"unicode/utf8"
)

type WsGetter interface {
	Ws() IWs
}

type ws struct {
}

func NewWs() *ws {
	return &ws{}
}

type IWs interface {
	GenerateConn(ws *websocket.Conn, client *ssh.Client, session *ssh.Session, channel ssh.Channel) error
}

func (w *ws) GenerateConn(ws *websocket.Conn, client *ssh.Client, session *ssh.Session, channel ssh.Channel) error {
	go func() {
		for {
			// 从websocket中读取数据
			_, p, err := ws.ReadMessage()
			if err != nil {
				return
			}
			// 将接收到的数据通过ssh channel通道写入
			_, err = channel.Write(p)
			if err != nil {
				return
			}
		}
	}()

	go func() {
		reader := bufio.NewReader(channel)
		var buf []byte

		t := time.NewTimer(time.Microsecond * 50)
		defer t.Stop()
		// 构建一个信道, 一端将数据远程主机的数据写入, 一段读取数据写入ws
		r := make(chan rune)
		go func() {
			defer client.Close()
			defer session.Close()
			for {
				x, size, err := reader.ReadRune()
				if err != nil {
					log.Logger.Error(err) //TODO control + D 会一直刷新
					ws.WriteMessage(1, []byte("\033[31m已经关闭连接!\033[0m"))
				}
				if size > 0 {
					r <- x
				}
			}
		}()

		for {
			select {
			case <-t.C:
				if len(buf) != 0 {
					err := ws.WriteMessage(websocket.TextMessage, buf)
					buf = []byte{}
					if err != nil {
						log.Logger.Error(err)
						return
					}
				}
				t.Reset(time.Microsecond * 50)
			case d := <-r:
				if d != utf8.RuneError {
					p := make([]byte, utf8.RuneLen(d))
					utf8.EncodeRune(p, d)
					buf = append(buf, p...)
				} else {
					buf = append(buf, []byte("@")...)
				}

			}
		}

	}()

	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error(err)
		}
	}()

	return nil
}

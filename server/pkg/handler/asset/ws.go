package asset

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/asset"
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

type WsMsg struct {
	Type int    `json:"type"`
	Msg  string `json:"msg"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
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
			var wsmsg WsMsg
			err = json.Unmarshal(p, &wsmsg)
			if err != nil {
				log.Logger.Error(err)
				return
			}
			// 将接收到的数据通过ssh channel通道写入
			//stdinPipe, err := session.StdinPipe()
			//_, err = stdinPipe.Write(p)
			switch wsmsg.Type {
			case 2:
				channel.Write([]byte(wsmsg.Msg))
				if err != nil {
					return
				}
			case 3:
				session.SendRequest("ping", true, nil)
				if err != nil {
					ws.WriteMessage(1, []byte("\033[31m已经关闭连接!\033[0m"))
					return
				}
			case 1:
				req := asset.TerminalWindow{
					Columns: uint32(wsmsg.Cols),
					Rows:    uint32(wsmsg.Rows),
					Width:   uint32(wsmsg.Cols * 8),
					Height:  uint32(wsmsg.Rows * 8),
				}
				_, err := session.SendRequest("window-change", false, ssh.Marshal(req))
				if err != nil {
					ws.WriteMessage(1, []byte("\033[31m已经关闭连接!\033[0m"))
					return
				}

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
					return
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

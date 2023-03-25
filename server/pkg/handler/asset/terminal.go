package asset

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/services"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type TerminalGetter interface {
	Terminal() ITerminal
}

type terminal struct {
	factory services.FactoryImp
}

func NewTerminal(f services.FactoryImp) *terminal {
	return &terminal{factory: f}
}

type ITerminal interface {
	GenerateClient(ctx context.Context, hostID int64, col, row int) (*ssh.Client, *ssh.Session, ssh.Channel, error)
}

func (t *terminal) GenerateClient(ctx context.Context, hostID int64, col, row int) (client *ssh.Client, session *ssh.Session, channel ssh.Channel, err error) {

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		config       ssh.Config
	)
	res, err := t.factory.Host().GetByHostId(ctx, hostID)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(res.Password))
	config = ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}
	clientConfig = &ssh.ClientConfig{
		User:    res.Username,
		Auth:    auth,
		Timeout: 5 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", res.Ip, res.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Logger.Error(err)
		return
	}
	session, channel, err = t.generateRequestTerminal(ctx, client, col, row)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}

func (t *terminal) generateRequestTerminal(ctx context.Context, client *ssh.Client, col, row int) (*ssh.Session, ssh.Channel, error) {
	session, err := client.NewSession()
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	channel, inRequests, err := client.OpenChannel("session", nil)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}
	// 处理无需返回数据的channel
	go func() {
		for req := range inRequests {
			if req.WantReply {
				req.Reply(false, nil)
			}
		}
	}()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	var modeList []byte

	for k, v := range modes {
		kv := struct {
			Key   byte
			Value uint32
		}{k, v}
		modeList = append(modeList, ssh.Marshal(kv)...)
	}
	modeList = append(modeList, 0)

	// 生成负载信息-需要传递的信息
	req := asset.PtyRequestMsg{
		Term:     "xterm",
		Columns:  uint32(col),
		Rows:     uint32(row),
		Width:    uint32(col) * 8,
		Height:   uint32(row) * 8,
		Modelist: string(modeList),
	}

	ok, err := channel.SendRequest("pty-req", true, ssh.Marshal(&req))
	if !ok || err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	ok, err = channel.SendRequest("shell", true, nil)
	if !ok || err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}

	return session, channel, nil
}

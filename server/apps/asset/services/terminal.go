package services

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/apps/asset/entity"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type TerminalGetter interface {
	Terminal() ITerminal
}

type terminal struct {
	host    IHost
	account IAccount
}

func NewTerminal(host IHost, account IAccount) ITerminal {
	return &terminal{host: host, account: account}
}

type ITerminal interface {
	GenerateClient(ctx context.Context, hostID uint64, accountId uint64, col, row int) (*ssh.Client, *ssh.Session, ssh.Channel)
}

func (t *terminal) GenerateClient(ctx context.Context, hostID uint64, accountId uint64, col, row int) (client *ssh.Client, session *ssh.Session, channel ssh.Channel) {

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		config       ssh.Config
	)
	host := t.host.GetByHostId(ctx, hostID)
	account := t.account.GetByAccountId(ctx, accountId)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(util.Decrypt(account.Password)))
	config = ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}
	clientConfig = &ssh.ClientConfig{
		User:    account.UserName,
		Auth:    auth,
		Timeout: 5 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", host.Ip, host.Port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	session, channel = t.generateRequestTerminal(ctx, client, col, row)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return
}

func (t *terminal) generateRequestTerminal(ctx context.Context, client *ssh.Client, col, row int) (*ssh.Session, ssh.Channel) {
	session, err := client.NewSession()
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	channel, inRequests, err := client.OpenChannel("session", nil)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
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
	req := entity.PtyRequestMsg{
		Term:     "xterm",
		Columns:  uint32(col),
		Rows:     uint32(row),
		Width:    uint32(col) * 8,
		Height:   uint32(row) * 8,
		Modelist: string(modeList),
	}

	ok, err := channel.SendRequest("pty-req", true, ssh.Marshal(&req))
	if !ok || err != nil {
		global.Logger.Error(err)
		return nil, nil
	}

	ok, err = channel.SendRequest("shell", true, nil)
	if !ok || err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	return session, channel
}

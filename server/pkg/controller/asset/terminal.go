package asset

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/services"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type TerminalGetter interface {
	Terminal() ITerminal
}

type terminal struct {
	factory services.IDbFactory
}

func NewTerminal(f services.IDbFactory) *terminal {
	return &terminal{f}
}

type ITerminal interface {
	GenerateClient(ctx context.Context, hostID int64) (*ssh.Client, error)
}

func (t *terminal) GenerateClient(ctx context.Context, hostID int64) (*ssh.Client, error) {

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		err          error
	)
	res, err := t.factory.Host().GetByHostId(ctx, hostID)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
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
		return nil, err
	}
	return client, nil
}

//func (t *terminal) GenerateRequestTerminal(ctx context.Context) (*ssh.Client, error) {
//
//}
